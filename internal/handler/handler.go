package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	type healthyResponse struct {
		Status string `json:"status"`
	}
	if err := json.NewEncoder(w).Encode(healthyResponse{"ok"}); err != nil {
		log.Fatalf("json.encode: %v", err)
		return
	}
}


func helloHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/healthz", http.StatusPermanentRedirect)
}
