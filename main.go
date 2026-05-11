package main

import (
	"log"
	"terrariadle-backend/internal/db"
	"terrariadle-backend/internal/jobs"
	"terrariadle-backend/internal/server"
	"terrariadle-backend/internal/store"
)

func main() {
	// Connect to the database
	db.Connect()
	defer db.DisconnectDB()

	// Initialize atomicstore
	store.InitializeStoreFromJson()

	// Sets up endpoints
	h := server.NewMux()

	// Starts the reset job
	go jobs.StartResetJob()

	// Starts the server
	if err := server.RunServer(":3000", h); err != nil {
		log.Fatal(err)
	}
}
