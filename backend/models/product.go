package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Product struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Price       string             `json:"price"`
	Images      []string           `json:"images"` // Store Base64 image strings
	CreatedBy   string             `json:"createdBy"` // Store admin ID
	CreatedAt   time.Time          `json:"createdAt"`
}
