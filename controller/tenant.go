package controller

import (
	"clinic-hub/models"
	"clinic-hub/services"
)

func CreateTenant(tenant models.TenantRequest) string {
	return services.CreateTenant(tenant)
}

func GetAllTenant() []models.Tenant {
	return services.GetAllTenant()
}
