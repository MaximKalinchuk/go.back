package goback

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/viper"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	fmt.Printf("\nServer starting on port: %s", viper.GetString("port"))

	server := s.httpServer.ListenAndServe()

	return server
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
