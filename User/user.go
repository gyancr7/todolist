package User

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name    string `json:"name"`
	EmailId string `json:"emailId"`
}


func saveNewUser(user *User, db *gorm.DB) {
	db.Create(user)
}