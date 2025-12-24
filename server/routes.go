package main

import (
	"github.com/go-chi/chi"
)

func addRoutes(r chi.Router) {
	r.Get("/healthz", healthHandler)

	r.Route("/hello", func(r chi.Router) {
		r.Use(SuperSafeAuth)

		r.Get("/", helloHandler)
	})
}
