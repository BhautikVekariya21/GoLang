package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/BhautikVekariya21/go/backend/routes"
	"github.com/BhautikVekariya21/go/backend/utils"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// loggingMiddleware is a simple middleware to log incoming requests
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r) // Call the next handler
		log.Printf("Method: %s, Path: %s, Duration: %s", r.Method, r.URL.Path, time.Since(start))
	})
}

// recoverMiddleware is a middleware to recover from panics and return a 500 status
func recoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				log.Printf("Recovered from panic: %v", err)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to MongoDB
	utils.ConnectDB() // Initialize the MongoDB client

	router := mux.NewRouter()
	routes.RegisterRoutes(router)

	// Apply middleware
	router.Use(loggingMiddleware)
	router.Use(recoverMiddleware)

	log.Printf("Starting server on port %s...", os.Getenv("PORT"))
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), router); err != nil {
		log.Fatal(err)
	}
}
