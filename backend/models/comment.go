package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Comment represents a user comment
type Comment struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Name    string             `bson:"name"`
	Email   string             `bson:"email"`
	Message string             `bson:"message"`
}
