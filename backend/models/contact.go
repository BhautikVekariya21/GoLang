package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Contact represents a user contact message
type Contact struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Name    string             `bson:"name"`
	Email   string             `bson:"email"`
	Message string             `bson:"message"`
}
