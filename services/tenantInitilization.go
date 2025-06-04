package services

import (
	"context"
	"fmt"
	"clinic-hub/database"
	"clinic-hub/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func TenantInitilization() {

	collection := database.Connection("fuzei", "tenant")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"name": "fuzei"}
	result := collection.FindOne(ctx, filter)
	var tenant models.Tenant
	err := result.Decode(&tenant)

	if err == mongo.ErrNoDocuments {
		newTenant := models.Tenant{
			Name:         "fuzei",
			AdminName:    "furkan",
			AdminSurname: "memiş",
			AdminEmail:   "furkan@fuzei.com",
			UUID:         "fuzei",
		}

		collection.InsertOne(ctx, newTenant)

		fmt.Println("Tenant initilization success " + newTenant.Name)
	} else {
		fmt.Println("Tenant already exist!")
	}

}
