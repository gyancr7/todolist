package Database

import (
	"github.com/jinzhu/gorm"
	"todolist/Task"
	"todolist/User"
)

func RunMigrations(db *gorm.DB) {
	db.Debug().DropTableIfExists(&User.User{})
	db.Debug().AutoMigrate(&User.User{})

	db.Debug().DropTableIfExists(&Task.Task{})
	db.Debug().AutoMigrate(&Task.Task{})
}
