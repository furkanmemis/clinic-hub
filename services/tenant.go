package services

import (
	"clinic-hub/database"
	"clinic-hub/models"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateTenant(tenant models.TenantRequest) string {
	collection := database.Connection("fuzei", "tenant")
	collectionTenantUser := database.Connection("fuzei", "tenantuser")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"name": tenant.Name}
	result := collection.FindOne(ctx, filter)

	var currentTenant models.Tenant
	err := result.Decode(&currentTenant)

	if err == mongo.ErrNoDocuments {

		newUuid := uuid.New().String()

		newTenant := models.Tenant{
			Name:         tenant.Name,
			AdminName:    tenant.AdminName,
			AdminSurname: tenant.AdminSurname,
			AdminEmail:   tenant.AdminEmail,
			UUID:         newUuid,
			Address:      tenant.Address,
			Phone:        tenant.Phone,
		}

		_, err := collection.InsertOne(ctx, newTenant)
		if err != nil {
			fmt.Println("Tenant creation failed:", err)
			return "Error: " + err.Error()
		}

		fmt.Println(tenant.Name + " Tenant created.")

		collectionUser := database.Connection(newUuid, "user")

		data := []byte(tenant.AdminPassword)
		hash := sha256.Sum256(data)

		newAdmin := models.User{
			Name:     tenant.AdminName,
			Surname:  tenant.AdminName,
			Email:    tenant.AdminEmail,
			Password: hex.EncodeToString(hash[:]),
			Role:     "manager",
			TenantId: newUuid,
		}

		newTenantUser := models.TenantUser{
			Email:    tenant.AdminEmail,
			Password: hex.EncodeToString(hash[:]),
			Role:     "manager",
			TenantId: newUuid,
		}

		collectionUser.InsertOne(ctx, newAdmin)
		collectionTenantUser.InsertOne(ctx, newTenantUser)
		RoleInitilization(newUuid)

		return "Tenant created: " + tenant.Name

	} else if err != nil {
		fmt.Println("Error:", err)
		return "Error: " + err.Error()
	}

	fmt.Println("Tenant already exists!")
	return tenant.Name + " already exists!"
}

func GetAllTenant() []models.Tenant {
	collection := database.Connection("fuzei", "tenant")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	var tenants []models.Tenant

	for cursor.Next(ctx) {
		var tenant models.Tenant
		if err := cursor.Decode(&tenant); err != nil {
			log.Fatal(err)
		}
		tenants = append(tenants, tenant)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	return tenants
}
