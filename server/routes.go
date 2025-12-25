package server

import (
	"net/http"

	"first-api/internal/handler"
	"first-api/internal/middleware"
	"github.com/go-chi/chi"
)

func addRoutes(r chi.Router) {
	r.Mount("/", http.FileServer(http.Dir("./static")))
	r.Get("/healthz", handler.HealthHandler)

	r.Route("/hello", func(r chi.Router) {
		r.Use(middleware.SuperSafeAuth)

		r.Get("/", handler.HelloHandler)
	})
}
