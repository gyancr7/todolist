package User

import (
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"net/http"
	"todolist/Utility"
)

func CreateNewUser(response http.ResponseWriter, request *http.Request, db *gorm.DB) {
	reqBody, _ := ioutil.ReadAll(request.Body)

	var user User

	Utility.UnmarshalInput(reqBody, &user)

	saveNewUser(&user, db)
}
