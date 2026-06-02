package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"terrariadle-backend/internal/domain"
)

func getUserID(r *http.Request) (string, bool) {
	id := r.URL.Query().Get("user_id")
	return id, id != ""
}

func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{"error": msg})
}

func handleError(w http.ResponseWriter, err error) {
	var appErr *domain.AppError
	if errors.As(err, &appErr) {
		switch appErr.Code {
		case domain.ErrUsrNotFound:
			writeError(w, http.StatusUnauthorized, appErr.Message)
		case domain.ErrNotFound:
			writeError(w, http.StatusNotFound, appErr.Message)
		case domain.ErrConflict:
			writeError(w, http.StatusConflict, appErr.Message)
		case domain.ErrInvalidInput:
			writeError(w, http.StatusBadRequest, appErr.Message)
		case domain.ErrInternal:
			writeError(w, http.StatusInternalServerError, "something went wrong")
		default:
			writeError(w, http.StatusInternalServerError, "something went wrong")
		}
		return
	}
	writeError(w, http.StatusInternalServerError, "something went wrong")
}
