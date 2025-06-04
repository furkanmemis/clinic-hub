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

func Login(email string, password string) []models.TenantUser {
	collection := database.Connection("fuzei", "tenantuser")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	data := []byte(password)
	hash := sha256.Sum256(data)
	hashedPassword := hex.EncodeToString(hash[:])

	fmt.Println("hashed password -> ", hashedPassword)
	fmt.Println("email -> ", email)

	filter := bson.M{"email": email, "password": hashedPassword}

	// projection := bson.M{"password": 0}
	// opts := options.Find().SetProjection(projection)

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		fmt.Println("Find error: ", err)
		return []models.TenantUser{}
	}
	defer cursor.Close(ctx)

	var tenantUsers []models.TenantUser
	if err := cursor.All(ctx, &tenantUsers); err != nil {
		fmt.Println("Cursor All error: ", err)
		return []models.TenantUser{}
	}

	return tenantUsers
}

func TenantLogin(tenantLoginRequest models.TenantLoginRequest) models.TenantLoginResponse {

	/*
		TODO
		response a tenant name i de ekle
	*/

	collection := database.Connection(tenantLoginRequest.TenantId, "user")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"email": tenantLoginRequest.Email, "tenantId": tenantLoginRequest.TenantId}

	result := collection.FindOne(ctx, filter)

	var user models.User

	err := result.Decode(&user)

	if err != mongo.ErrNoDocuments {

		token, err := GenerateJWT(user.Email, user.TenantId)

		if err != nil {
			fmt.Println("Token create error: ", err)
			var lgr models.TenantLoginResponse
			return lgr
		}

		loginResponse := models.TenantLoginResponse{
			Name:       user.Name,
			Surname:    user.Surname,
			TenantName: user.TenantId,
			Role:       user.Role,
			Token:      token,
			Email:      user.Email,
		}

		return loginResponse

	} else {
		var lgr models.TenantLoginResponse
		return lgr
	}

}
