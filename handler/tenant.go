package handler

import (
	"encoding/json"
	"clinic-hub/controller"
	"clinic-hub/models"
	"net/http"
)

func CreateTenantHandler(w http.ResponseWriter, r *http.Request) {

	var tenant models.TenantRequest

	err := json.NewDecoder(r.Body).Decode(&tenant)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	password := tenant.AdminPassword

	msg := controller.CreateTenant(tenant.TenantInformation, password)

	rsp := map[string]interface{}{
		"message": "Success",
		"tenant":  msg,
	}

	json.NewEncoder(w).Encode(rsp)

}

var TenantHandlerMap = map[string]func(http.ResponseWriter, *http.Request){
	"/tenant/create": CreateTenantHandler,
}
