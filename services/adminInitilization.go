package services

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"clinic-hub/database"
	"clinic-hub/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func AdminInitilization() {
	collection := database.Connection("fuzei", "user")
	collectionTenant := database.Connection("fuzei", "tenant")
	collectionTenantUser := database.Connection("fuzei", "tenantuser")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	filter := bson.M{"role": "admin"}
	result := collection.FindOne(ctx, filter)

	var admin models.User
	err := result.Decode(&admin)

	if err == mongo.ErrNoDocuments {

		filter := bson.M{"name": "fuzei"}
		resultTenant := collectionTenant.FindOne(ctx, filter)
		var currentTenant models.Tenant

		err := resultTenant.Decode(&currentTenant)
		if err != nil {
			fmt.Println("Tenant not found:", err)
			return
		}

		data := []byte("furkanzeynep")
		hash := sha256.Sum256(data)

		newAdmin := models.User{
			Name:     "Furkan",
			Surname:  "Memiş",
			Email:    "furkan@fuzei.com",
			Password: hex.EncodeToString(hash[:]),
			Role:     "admin",
			TenantId: currentTenant.UUID,
		}

		newTenantUser := models.TenantUser{
			Email:    "furkan@fuzei.com",
			Password: hex.EncodeToString(hash[:]),
			Role:     "admin",
			TenantId: currentTenant.UUID,
		}

		collection.InsertOne(ctx, newAdmin)
		collectionTenantUser.InsertOne(ctx, newTenantUser)

		fmt.Println("Admin created.")
	} else {
		fmt.Println("Admin already exist!")
		return
	}

}
