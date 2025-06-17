package handler

import (
	"clinic-hub/controller"
	middleware "clinic-hub/middlewares"
	"clinic-hub/models"
	"encoding/json"
	"net/http"
)

func CreateDepartmentHandler(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(middleware.UserKey).(map[string]string)
	if !ok || user == nil {
		http.Error(w, "User not found in context", http.StatusUnauthorized)
		return
	}
	tenantID := user["tenantId"]

	var departmentModel models.Department

	err := json.NewDecoder(r.Body).Decode(&departmentModel)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	msg := controller.CreateDepartment(departmentModel, tenantID)
	rsp := map[string]interface{}{
		"message":    "Success",
		"department": msg,
	}

	json.NewEncoder(w).Encode(rsp)
}

var DeparmentHandler = map[string]func(http.ResponseWriter, *http.Request){
	"/department/create": CreateDepartmentHandler,
}
