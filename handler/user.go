package handler

import (
	"encoding/json"
	"clinic-hub/controller"
	"clinic-hub/models"
	"net/http"
	"strings"
)

func GetAllUserHandler(w http.ResponseWriter, r *http.Request) {
	users := controller.GetAllUsers()
	rsp := map[string]interface{}{
		"users":   users,
		"message": "Success",
	}
	json.NewEncoder(w).Encode(rsp)
}

func GetUserByIdHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	parts := strings.Split(path, "/")

	if len(parts) < 4 {
		http.Error(w, "Id not provided", http.StatusBadRequest)
		return
	}

	id := parts[3]

	user := controller.GetUserById(id)

	rsp := map[string]interface{}{
		"message": "Succes",
		"user":    user,
	}

	json.NewEncoder(w).Encode(rsp)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	msg := controller.CreateUser(user)

	rsp := map[string]interface{}{
		"message": "Success",
		"user":    msg,
	}

	json.NewEncoder(w).Encode(rsp)
}

var UserHandlerMap = map[string]func(http.ResponseWriter, *http.Request){
	"/user/get-all": GetAllUserHandler,
	"/user/get-id/": GetUserByIdHandler,
	"/user/create":  CreateUser,
}
