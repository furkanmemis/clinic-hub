package handler

import (
	"encoding/json"
	"clinic-hub/controller"
	"clinic-hub/models"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	var loginRequest models.LoginRequest

	err := json.NewDecoder(r.Body).Decode(&loginRequest)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	tenantUsers := controller.Login(loginRequest)

	rsp := map[string]interface{}{
		"message":    "Success",
		"tenantUser": tenantUsers,
	}

	json.NewEncoder(w).Encode(rsp)

}

func TenantLoginHandler(w http.ResponseWriter, r *http.Request) {
	var tenantLoginRequest models.TenantLoginRequest

	err := json.NewDecoder(r.Body).Decode(&tenantLoginRequest)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	tenantLoginResponse := controller.TenantLogin(tenantLoginRequest)

	rsp := map[string]interface{}{
		"message": "Success",
		"user":    tenantLoginResponse,
	}

	json.NewEncoder(w).Encode(rsp)
}

var AuthenticationHandlerMap = map[string]func(http.ResponseWriter, *http.Request){
	"/auth/login":        LoginHandler,
	"/auth/tenant-login": TenantLoginHandler,
}
