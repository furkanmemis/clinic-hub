package handler

import (
	"clinic-hub/controller"
	"clinic-hub/models"
	"encoding/json"
	"net/http"
)

func CreateTenantHandler(w http.ResponseWriter, r *http.Request) {

	var tenant models.TenantRequest

	err := json.NewDecoder(r.Body).Decode(&tenant)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	msg := controller.CreateTenant(tenant)

	rsp := map[string]interface{}{
		"message": "Success",
		"tenant":  msg,
	}

	json.NewEncoder(w).Encode(rsp)

}

func GetAllTenantHandler(w http.ResponseWriter, r *http.Request) {
	tenants := controller.GetAllTenant()
	json.NewEncoder(w).Encode(tenants)
}

var TenantHandlerMap = map[string]func(http.ResponseWriter, *http.Request){
	"/tenant/create":  CreateTenantHandler,
	"/tenant/get-all": GetAllTenantHandler,
}
