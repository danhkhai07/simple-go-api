package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
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


func HelloHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/healthz", http.StatusPermanentRedirect)
}
