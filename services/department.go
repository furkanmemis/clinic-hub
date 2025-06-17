package services

import (
	"clinic-hub/database"
	"clinic-hub/models"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateDepartment(department models.Department, tenantId string) string {

	collection := database.Connection(tenantId, "department")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	result, err := collection.InsertOne(ctx, department)
	if err != nil {
		fmt.Println("Create user error: ", err)
		return ""
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()
	return id
}
