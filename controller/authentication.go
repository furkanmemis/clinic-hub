package controller

import (
	"clinic-hub/models"
	"clinic-hub/services"
)

func Login(loginRequest models.LoginRequest) []models.TenantUser {
	return services.Login(loginRequest.Email, loginRequest.Password)
}

func TenantLogin(tenantLoginRequest models.TenantLoginRequest) models.TenantLoginResponse {
	return services.TenantLogin(tenantLoginRequest)
}
