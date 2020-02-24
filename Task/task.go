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

func getAllTasks(userId string, db *gorm.DB) []Task{

	rows, err :=  db.Table("tasks").Where(&Task{UserId: userId}).Rows()

	return 	iterateTaskRows(rows, err, db)
}

func deleteTasks(userId string, taskId string, db *gorm.DB) string{

	db.Table("tasks").Where("id=? and user_id=?", taskId, userId).Delete(&Task{})

	return fmt.Sprintf("Task: %s deleted for the User %s", taskId, userId)
}

func saveNewTask(task *Task, db *gorm.DB) {
	db.Create(&task)
}

func iterateTaskRows(taskRows *sql.Rows, err error, db *gorm.DB) []Task{
	var tasks []Task

	if err != nil {
		fmt.Fprintf(os.Stderr, "Fprintf: %v\n", err)
	}

	defer taskRows.Close()

	for taskRows.Next() {
		var task Task

		err := db.ScanRows(taskRows, &task)

		if err != nil {
			log.Fatal(err)
		}

		tasks = append(tasks, task)
	}

	return tasks
}