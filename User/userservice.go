package User

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"net/http"
)

func CreateNewUser(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	fmt.Fprintf(w, "%+v", string(reqBody))

	var user User

	err := json.Unmarshal([]byte(string(reqBody)), &user)

	if err != nil {
		fmt.Printf("Error Unmarshalling")
	}

	saveNewUser(&user, db)
}
