package main

import (
	"log"
	"todolist/Database"
	"todolist/Server"
)

func main() {

	db, err := Database.InitDbConnection()

	if err != nil {
		log.Panic(err)
	}

	log.Println("Database Connection Established")

	defer db.Close()

	Database.RunMigrations(db)

	Server.HandleRequests(db)
	}