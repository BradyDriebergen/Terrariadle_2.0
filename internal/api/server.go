package api

import (
	"context"
	"net/http"
	"terrariadle-backend/internal/domain"
	"terrariadle-backend/internal/services"
)

type Server struct {
	httpServer  *http.Server
	dailySlash  services.DailySlashService
	connections services.ConnectionsService
	guessTheNpc services.GuessTheNpcService
	hangman     services.HangmanService
	sseServer   services.SseStreamService
	broker      domain.GuessCountBroker
}

func NewServer(
	addr string,
	dailySlash services.DailySlashService,
	connections services.ConnectionsService,
	guessTheNpc services.GuessTheNpcService,
	hangman services.HangmanService,
	sseServer services.SseStreamService,
	broker domain.GuessCountBroker,
) *Server {

	s := &Server{
		dailySlash:  dailySlash,
		connections: connections,
		guessTheNpc: guessTheNpc,
		hangman:     hangman,
		sseServer:   sseServer,
		broker:      broker,
	}

	s.httpServer = &http.Server{
		Addr:    addr,
		Handler: s.newMux(),
	}

	return s
}

func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
