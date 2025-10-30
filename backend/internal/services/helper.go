package services

import (
	"github.com/google/uuid"
)

// Checks if a UUID is valid
func isValidUUID(id string) bool {
	if id == "" {
		return false
	}
	_, err := uuid.Parse(id)
	return err == nil
}
