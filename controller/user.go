package controller

import (
	"clinic-hub/models"
	"clinic-hub/services"
)

func GetAllUsers() []models.User {
	return services.GetAllUsers()
}

func GetUserById(id string) models.User {
	return services.GetUserById(id)
}

func CreateUser(user models.User, tenantId string) string {
	return services.CreateUser(tenantId, user)
}
