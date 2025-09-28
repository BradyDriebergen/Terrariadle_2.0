package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func RunServer(addr string, h http.Handler) error {
	srv := &http.Server{
		Addr:         addr,
		Handler:      h,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Creates goRoutine for the server
	go func() {
		log.Printf("server: listening on %s", addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server: listen: %v", err)
		}
	}()

	// Recieves signal to close server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	log.Println("server: shutdown signal received")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Tries to gracefully shut down server
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("server: graceful shutdown failed: %v", err)

		// Force closes the server
		if cerr := srv.Close(); cerr != nil {
			log.Printf("server: forced close error: %v", cerr)
		}
	}

	log.Println("server: exited cleanly")
	return nil
}
