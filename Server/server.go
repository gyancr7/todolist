package Server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"todolist/Task"
	"todolist/User"
)

var dbConn *gorm.DB

func HandleRequests(db *gorm.DB) {

	dbConn = db

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/v1", homePage)

	router.HandleFunc("/v1/addUser", addUser)

	router.HandleFunc("/v1/user/assign_task", assignTask)

	router.HandleFunc("/v1/user/{id}/getAllTasks", getAllTasks)

	router.HandleFunc("/v1/user/{user_id}/delete_task/{task_id}", deleteTask).Methods("DELETE")

	router.HandleFunc("/v1/user/{user_id}/update_task/{task_id}", updateTask).Methods("PATCH")

	log.Fatal(http.ListenAndServe(":8081", router))
}

func homePage(response http.ResponseWriter, request *http.Request){
	fmt.Fprintf(response, "Welcome to the TodoList App!")
}

func addUser(response http.ResponseWriter, request *http.Request)  {
	User.CreateNewUser(response, request, dbConn)
}

func assignTask(response http.ResponseWriter, request *http.Request) {
	Task.AssignNewTask(response, request, dbConn)
}

func getAllTasks(response http.ResponseWriter, request *http.Request) {
	Task.GetAllTasks(response, request, dbConn)
}

func deleteTask(response http.ResponseWriter, request *http.Request) {
	Task.DeleteTask(response, request, dbConn)
}

func updateTask(response http.ResponseWriter, request *http.Request) {
	Task.UpdateTask(response, request, dbConn)
}

