package presentation

import (
	"crypto/tls"
	"mock-shipping-provider/business"
	"net"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

type Presenter struct {
	shippingService business.Shipping
}

type Dependency struct {
	ShippingService business.Shipping
}

type Config struct {
	Hostname  string
	Port      string
	TLSConfig *tls.Config

	Dependency Dependency
}

func NewHttpServer(config Config) (*http.Server, error) {
	router := chi.NewRouter()
	presenter := Presenter{
		shippingService: config.Dependency.ShippingService,
	}

	router.Post("/estimate", presenter.EstimateHandler)
	router.Post("/order", presenter.OrderHandler)
	router.Get("/status-history", presenter.StatusHistoryHandler)

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
