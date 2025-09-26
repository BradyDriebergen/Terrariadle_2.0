package main

import (
	"terrariadle-backend/internal/db"
)

func main() {
	// Connect once at startup
	db.Connect()

	// Get collection
	collection := db.GetCollection("backend-tests", "user-guesses")

}
