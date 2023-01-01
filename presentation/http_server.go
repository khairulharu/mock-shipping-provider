package presentation

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

type Dependency struct{}

type Config struct {
	Hostname  string
	Port      string
	TLSConfig *tls.Config

	Dependency Dependency
}

func NewHttpServer(config Config) (*http.Server, error) {
	router := chi.NewRouter()

	router.Post("/estimate", config.Dependency.EstimateHandler)
	router.Post("/order", config.Dependency.OrderHandler)
	router.Get("/status-history", config.Dependency.StatusHistoryHandler)

	server := &http.Server{
		Addr:              net.JoinHostPort(config.Hostname, config.Port),
		Handler:           router,
		TLSConfig:         config.TLSConfig,
		ReadTimeout:       time.Minute,
		ReadHeaderTimeout: time.Minute,
		WriteTimeout:      time.Minute,
		IdleTimeout:       time.Minute,
	}

	return server, nil
}
