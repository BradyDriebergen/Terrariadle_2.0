package main

import (
	"context"
	"log"
	"terrariadle-backend/internal/db"
	"terrariadle-backend/internal/jobs"
	"terrariadle-backend/internal/routes"
	"terrariadle-backend/internal/server"
	"terrariadle-backend/internal/types"
)

func main() {
	// Setup
	db.Connect()
	r := routes.SetupRouter()
	gameData := &types.GameData{}

	// Starts the reset job
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go jobs.StartResetJob(ctx, gameData)

	// Starts the server
	if err := server.RunServer(":3000", r); err != nil {
		log.Fatal(err)
	}
}
