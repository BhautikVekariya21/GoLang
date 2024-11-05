package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/BhautikVekariya21/go/backend/models"
	"github.com/BhautikVekariya21/go/backend/services"
)

// CommentHandler handles the creation of new comments
func CommentHandler(w http.ResponseWriter, r *http.Request) {
	var comment models.Comment

	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := services.CreateComment(comment); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(comment)
}

// GetCommentsHandler retrieves all comments
func GetCommentsHandler(w http.ResponseWriter, r *http.Request) {
	comments, err := services.GetAllComments()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(comments)
}
