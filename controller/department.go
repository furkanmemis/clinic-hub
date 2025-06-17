package controller

import (
	"clinic-hub/models"
	"clinic-hub/services"
)

func CreateDepartment(department models.Department, tenantId string) string {
	return services.CreateDepartment(department, tenantId)
}
