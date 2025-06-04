package controller

import (
	"clinic-hub/models"
	"clinic-hub/services"
)

func CreateTenant(tenant models.Tenant, password string) string {
	return services.CreateTenant(tenant, password)
}
