package main

import (
	"log"
	"terrariadle-backend/internal/db"
	"terrariadle-backend/internal/jobs"
	"terrariadle-backend/internal/server"
	"terrariadle-backend/internal/types"
	"terrariadle-backend/internal/utils"
)

func main() {
	// Setup
	db.Connect()
	r := server.SetupRouter()
	mStore := utils.NewMemoryStore[types.GameData]()

	// Starts the reset job
	go jobs.StartResetJob(mStore)

	// Starts the server
	if err := server.RunServer(":3000", r); err != nil {
		log.Fatal(err)
	}
}
