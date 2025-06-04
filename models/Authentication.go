package models

type LoginRequest struct {
	Email    string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
}

type TenantLoginRequest struct {
	Email    string `bson:"email" json:"email"`
	TenantId string `bson:"tenantId" json:"tenantId"`
}

type TenantLoginResponse struct {
	Email      string `bson:"email" json:"email"`
	TenantName string `bson:"tenantName" json:"tenantName"`
	Role       string `bson:"role" json:"role"`
	Name       string `bson:"name" json:"name"`
	Surname    string `bson:"surname" json:"surname"`
	Token      string `bson:"token" json:"token"`
}
