package main

import (
	"log"
	"terrariadle-backend/internal/db"
	"terrariadle-backend/internal/jobs"
	"terrariadle-backend/internal/server"
	"terrariadle-backend/internal/utils/cache"
)

func main() {
	// Setup
	db.Connect()
	r := server.SetupRouter()
	cache.NewGameStore()
	cache.NewPuzzleStore()

	// Starts the reset job
	go jobs.StartResetJob()

	// Starts the server
	if err := server.RunServer(":3000", r); err != nil {
		log.Fatal(err)
	}
}
