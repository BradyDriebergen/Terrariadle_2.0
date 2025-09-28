package main

import (
	"log"
	"terrariadle-backend/internal/routes"
	"terrariadle-backend/internal/server"
)

func main() {
	// Connect once at startup
	// db.Connect()

	r := routes.SetupRouter()

	if err := server.RunServer(":3000", r); err != nil {
		log.Fatal(err)
	}
}
