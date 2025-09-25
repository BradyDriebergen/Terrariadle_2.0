package main

import (
	"context"
	"fmt"
	"log"

	"terrariadle-backend/internal/db"
	"terrariadle-backend/internal/models"
)

func main() {
	// Connect once at startup
	db.Connect()

	// Get collection
	collection := db.GetCollection("backend-tests", "user-guesses")

	guessDoc := models.GuessDocument{
		UserID: "user123",
		Games: []models.Game{
			{
				GameType: "daily",
				Guesses:  []string{"Iron Sword", "Gold Pickaxe", "Diamond Helmet"},
				HasWon:   true,
				Position: 3,
				Extra: map[string]interface{}{
					"completedAt": "2025-09-24T10:30:00Z",
					"hints":       2,
				},
			},
			{
				GameType: "unlimited",
				Guesses:  []string{"Wooden Bow"},
				HasWon:   false,
				Position: -1,
				Extra: map[string]interface{}{
					"startedAt": "2025-09-24T11:00:00Z",
				},
			},
		},
	}

	// Example insert
	res, err := collection.InsertOne(context.Background(), guessDoc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted document ID:", res.InsertedID)

	err = collection.Drop(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Collection dropped successfully")
}
