package services

import "terrariadle-backend/internal/domain"

type ConnectionsInitData struct {
	Attepts          int
	Finished         bool
	Options          []domain.ConnectionOption
	SolvedCategories []SolvedCategory
}

type SolvedCategory struct {
	Name    string
	Options []string
}
