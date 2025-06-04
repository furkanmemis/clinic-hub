package services

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"clinic-hub/database"
	"clinic-hub/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var AllUser []models.User

func GetAllUsers() []models.User {
	return AllUser
}

func GetUserById(id string) models.User {
	for _, u := range AllUser {
		if u.ID.String() == id {
			return u
		}
	}
	return models.User{}
}

func CreateUser(databaseName string, user models.User) string {
	collection := database.Connection(databaseName, "user")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	data := []byte(user.Password)
	hash := sha256.Sum256(data)
	user.Password = hex.EncodeToString(hash[:])
	user.Role = "user"

	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		fmt.Println("Create user error: ", err)
		return ""
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()
	fmt.Println(user.Name + " " + user.Surname + " created.")

	return id
}
