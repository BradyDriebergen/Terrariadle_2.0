package api

import (
	"encoding/json"
	"net/http"
	"terrariadle-backend/internal/utils"
	"time"
)

func CheckHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func GetRemainingTime(w http.ResponseWriter, r *http.Request) {
	remaining := utils.TimeUntilNextMidnight(time.Now())

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"remainingSeconds": int(remaining.Seconds()),
	})
}
