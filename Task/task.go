package Task

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"os"
	)

type Task struct {
	gorm.Model
	Name      	string 		`json:"name"`
	Description string		`json:"description"`
	UserId    	string 		`json:"userId"`
	Estimate  	int    		`json:"estimate"`
	Completed 	bool   		`json:"completed"`
}

func getAllTasks(userId string, db *gorm.DB) []string{
	return 	iterateTaskRows(db.Table("tasks").Where(&Task{UserId: userId}).Rows())
}

func deleteTasks(userId string, taskId string, db *gorm.DB) string{
	var tasks = iterateTaskRows(db.Table("tasks").Where("id=? and user_id=?", taskId, userId).Rows())

	if len(tasks) == 0 {
		return fmt.Sprintf("Task: %s does not exist for the User %s", taskId, userId)
	}

	db.Table("tasks").Delete(&tasks)

	return fmt.Sprintf("Task: %s deleted for the User %s", taskId, userId)
}

func saveNewTask(task *Task, db *gorm.DB) {
	db.Create(&task)
}

func iterateTaskRows(taskRows *sql.Rows, err error) []string{
	var tasks []string

	if err != nil {
		fmt.Fprintf(os.Stderr, "Fprintf: %v\n", err)
	}

	defer taskRows.Close()

	for taskRows.Next() {
		var userId string

		err := taskRows.Scan(&userId)

		//fmt.Println(task)

		if err != nil {
			log.Fatal(err)
		}

		tasks = append(tasks, userId)
	}

	return tasks
}