package server

import (
	"net/http"
	"terrariadle-backend/internal/utils"
)

// Checks the status of the backend
func healthHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

// Gets the remaining time in the day in seconds
func getRemainingTime(w http.ResponseWriter, r *http.Request) {
	data := map[string]int64{
		"remaining": int64(utils.TimeUntilNextMidnightFromNow().Seconds()),
	}

	writeJSON(w, http.StatusOK, data)
}
