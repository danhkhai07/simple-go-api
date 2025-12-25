package handler

import (
	"encoding/json"
	"log"
	"os"
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
	w.Header().Set("Content-Type", "text/html")
	
	html, err := os.Open("./protected/hello.html")
	defer html.Close()
	if err != nil {
		http.Error(w, "file not found", http.StatusNotFound)
		return
	}

	stat, err := html.Stat()
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	http.ServeContent(
		w,
		r,
		stat.Name(),
		stat.ModTime(),
		html,
	)
}
