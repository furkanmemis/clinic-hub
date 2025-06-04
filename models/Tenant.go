package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tenant struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UUID         string             `bson:"UUID"`
	Name         string             `bson:"name" json:"name"`
	AdminName    string             `bson:"adminName" json:"adminName"`
	AdminSurname string             `bson:"adminSurname" json:"adminSurname"`
	AdminEmail   string             `bson:"adminEmail" json:"adminEmail"`
}

type TenantRequest struct {
	TenantInformation Tenant `json:"tenantInformation"`
	AdminPassword     string `json:"adminPassword"`
}
