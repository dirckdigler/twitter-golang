package routers

import (
	"encoding/json"
	"net/http"

	"github.com/dirckdigler/twitter/models"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error in data entry"+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "Email is mandatory", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "Password should containd minimun 6 characters", 400)
		return
	}
	// _, encontrado, _ := bd.CheckUserAlreadyExist(t.Email)
	// if encontrado == true {
	// 	http.Error(w, "There is already a registered user with that email", 400)
	// 	return
	// }

	// _, status, err := bd.InsertRegister(t)
	// if err != nil {
	// 	http.Error(w, "An error occurred while trying to register user "+err.Error(), 400)
	// 	return
	// }

	// if status == false {
	// 	http.Error(w, "Unable to insert user record", 400)
	// 	return
	// }

	w.WriteHeader(http.StatusCreated)
}
