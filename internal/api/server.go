package api

import (
	"context"
	"net"
	"net/http"
	"terrariadle-backend/internal/domain"
	"terrariadle-backend/internal/services"
)

type Server struct {
	httpServer  *http.Server
	cancelBase  context.CancelFunc
	dailySlash  services.DailySlashService
	connections services.ConnectionsService
	guessTheNpc services.GuessTheNpcService
	hangman     services.HangmanService
	terraTrivia services.TerraTrivia
	sseServer   services.SseStreamService
	broker      domain.GuessCountBroker
}

func NewServer(
	ctx context.Context,
	addr string,
	dailySlash services.DailySlashService,
	connections services.ConnectionsService,
	guessTheNpc services.GuessTheNpcService,
	hangman services.HangmanService,
	terraTrivia services.TerraTrivia,
	sseServer services.SseStreamService,
	broker domain.GuessCountBroker,
) *Server {
	baseCtx, cancelBase := context.WithCancel(ctx)

	s := &Server{
		cancelBase:  cancelBase,
		dailySlash:  dailySlash,
		connections: connections,
		guessTheNpc: guessTheNpc,
		hangman:     hangman,
		terraTrivia: terraTrivia,
		sseServer:   sseServer,
		broker:      broker,
	}

	s.httpServer = &http.Server{
		Addr:        addr,
		Handler:     s.newMux(),
		BaseContext: func(_ net.Listener) context.Context { return baseCtx },
	}

	return s
}

func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	s.cancelBase()
	return s.httpServer.Shutdown(ctx)
}
