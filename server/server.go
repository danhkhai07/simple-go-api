package server

import (
	"net/http"
	"log"
	"context"
	"os"
	"os/signal"
	"time"

	"first-api/config"
	"github.com/go-chi/chi"
)

func NewServer (
	logger *config.Logger,
	cfg *config.Config,
) http.Handler {
	r := chi.NewRouter()

	addRoutes(r)

	var handler http.Handler = r
	//global midware

	return handler
}

func Run(
	logger *config.Logger,
	cfg *config.Config,
) {
	// Build handler
	handler := NewServer(logger, cfg)

	// Setup HTTP server
	svr := &http.Server {
		Addr: cfg.Host + ":" + cfg.Port,
		Handler: handler,
	}
		
	// Channel to receive OS interrupt signal
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	go func() {
		log.Printf("Starting server on %s\n", svr.Addr)
		if err := svr.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-stop // Block until receive interrupt signal

	// Create context to gracefully shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Printf("Shutting down server...\n")
	if err := svr.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v\n", err)
	}

	log.Printf("Server exiting...\n")
}
