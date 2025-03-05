package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username     string             `json:"username"`
	Email    string             `json:"email"`
	MobileNumber   string             `bson:"mobileNumber" json:"mobileNumber"` // Mobile number
	Password string             `json:"password"`
	Role     string             `json:"role"` // "user", "admin", "superadmin"
}
