package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name     string             `bson:"name" json:"name"`
	Surname  string             `bson:"surname" json:"surname"`
	Password string             `bson:"password" json:"password"`
	Email    string             `bson:"email" json:"email"`
	Role     string             `bson:"role" json:"role"`
	TenantId string             `bson:"tenantId" json:"tenantId"`
}

type TenantUser struct {
	Password string `bson:"password" json:"password"`
	Email    string `bson:"email" json:"email"`
	Role     string `bson:"role" json:"role"`
	TenantId string `bson:"tenantId" json:"tenantId"`
}
