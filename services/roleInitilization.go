package services

import (
	"clinic-hub/database"
	"clinic-hub/models"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var roleArray = [4]string{"admin", "manager", "doctor", "visitor"}

func RoleInitilization(tenantId string) {

	collection := database.Connection(tenantId, "role")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	for _, role := range roleArray {

		filter := bson.M{"name": role}
		result := collection.FindOne(ctx, filter)

		var roleModel models.Role
		err := result.Decode(&roleModel)

		if err == mongo.ErrNoDocuments {
			newRole := models.Role{Name: role}
			collection.InsertOne(ctx, newRole)
		} else {
			fmt.Printf("%s role already exist for", role)
		}

	}

}
