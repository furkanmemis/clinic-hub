package handler

import (
	"clinic-hub/controller"
	middleware "clinic-hub/middlewares"
	"clinic-hub/models"
	"encoding/json"
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

	user, ok := r.Context().Value(middleware.UserKey).(map[string]string)
	if !ok || user == nil {
		http.Error(w, "User not found in context", http.StatusUnauthorized)
		return
	}

	//email := user["email"]
	tenantID := user["tenantId"]

	var userModel models.User

	err := json.NewDecoder(r.Body).Decode(&userModel)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	msg := controller.CreateUser(userModel, tenantID)

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
