package handlers

import (
	"encoding/json"
	"net/http"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := map[string]string{"status": "ok"}
	json.NewEncoder(w).Encode(resp)
}
