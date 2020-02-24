package Database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const(
	host = "127.0.0.1"
	user = "root"
	port = 23306
	database = "todolist"
	password = 123
)

func InitDbConnection() (*gorm.DB, error)  {
	var dbConnString = fmt.Sprintf("%s:%d@tcp(%s:%d)/%s?charset=utf8&parseTime=True", user, password, host, port, database)

	db, err := gorm.Open("mysql", dbConnString)

	return db, err
}



