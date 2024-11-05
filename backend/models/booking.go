package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Booking struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	Name              string             `bson:"name"`
	Email             string             `bson:"email"`
	EventDate         string             `bson:"event_date"`
	GuestCount        int                `bson:"guest_count"`
	DecorationCharges int                `bson:"decoration_charges"`
	LightBill         int                `bson:"light_bill"`
}
