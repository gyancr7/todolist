package Task

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"net/http"
	"todolist/Utility"
)

func AssignNewTask(response http.ResponseWriter, request *http.Request, db *gorm.DB) {
	reqBody, _ := ioutil.ReadAll(request.Body)

	var task Task

	Utility.UnmarshalInput(reqBody, &task)

	saveNewTask(&task, db)
}

func GetAllTasks(response http.ResponseWriter, request *http.Request, db *gorm.DB) {
	vars := mux.Vars(request)
	userId := vars["id"]

	getAllTasks(userId, db)
	Utility.MarshalOutput(getAllTasks(userId, db), response)
}

func DeleteTask(response http.ResponseWriter, request *http.Request, db *gorm.DB) {
	vars := mux.Vars(request)
	userId := vars["user_id"]
	taskId := vars["task_id"]

	Utility.MarshalOutput(deleteTasks(userId, taskId, db), response)
}
