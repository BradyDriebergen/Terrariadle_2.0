package services

type ConnectionsInitData struct {
	Attepts          int
	Finished         bool
	Options          []string
	SolvedCategories []SolvedCategory
}

type SolvedCategory struct {
	Name    string
	Options []string
}
