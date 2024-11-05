package controllers

import (
	"context"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/BhautikVekariya21/go/backend/models" // Full import path
	"github.com/BhautikVekariya21/go/backend/utils"  // Full import path
)

var contactCollection *mongo.Collection = utils.Client.Database("partyplot").Collection("contacts")

func CreateContact(w http.ResponseWriter, r *http.Request) {
	var contact models.Contact
	err := json.NewDecoder(r.Body).Decode(&contact)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insert the contact into the database
	_, err = contactCollection.InsertOne(context.TODO(), contact)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]string{"message": "Message sent successfully!"}
	json.NewEncoder(w).Encode(response)
}
