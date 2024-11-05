package services

import (
	"context"

	"github.com/BhautikVekariya21/go/backend/models"
	"github.com/BhautikVekariya21/go/backend/utils"
	"go.mongodb.org/mongo-driver/bson"
)

// CreateComment saves a new comment in the database
func CreateComment(comment models.Comment) error {
	collection := utils.Client.Database("partyplot").Collection("comments") // Ensure the DB name is correct
	_, err := collection.InsertOne(context.TODO(), comment)
	return err
}

// GetAllComments retrieves all comments from the database
func GetAllComments() ([]models.Comment, error) {
	var comments []models.Comment
	collection := utils.Client.Database("partyplot").Collection("comments")

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var comment models.Comment
		if err = cursor.Decode(&comment); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

// CreateContact saves a new contact message in the database
func CreateContact(contact models.Contact) error {
	collection := utils.Client.Database("partyplot").Collection("contacts") // Ensure the DB name is correct
	_, err := collection.InsertOne(context.TODO(), contact)
	return err
}

// GetAllContacts retrieves all contact messages from the database
func GetAllContacts() ([]models.Contact, error) {
	var contacts []models.Contact
	collection := utils.Client.Database("partyplot").Collection("contacts")

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var contact models.Contact
		if err = cursor.Decode(&contact); err != nil {
			return nil, err
		}
		contacts = append(contacts, contact)
	}
	return contacts, nil
}

// CreateBooking saves a new booking in the database
func CreateBooking(booking models.Booking) error {
	collection := utils.Client.Database("partyplot").Collection("bookings") // Ensure the DB name is correct
	_, err := collection.InsertOne(context.TODO(), booking)
	return err
}

// GetAllBookings retrieves all bookings from the database
func GetAllBookings() ([]models.Booking, error) {
	var bookings []models.Booking
	collection := utils.Client.Database("partyplot").Collection("bookings") // Ensure the DB name is correct

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var booking models.Booking
		if err = cursor.Decode(&booking); err != nil {
			return nil, err
		}
		bookings = append(bookings, booking)
	}
	return bookings, nil
}
