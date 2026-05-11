package jsonreader

import "fmt"

// Npc struct that gets pulled from JSON files

type Npc struct {
	ID      int      `json:"id"`
	NPC     string   `json:"npc"`
	NPCPath string   `json:"npc-path"`
	Quotes  []string `json:"quotes"`
	Names   []string `json:"names"`
}

type SearchNpcResult struct {
	NpcId int    `json:"id"`
	Name  string `json:"name"`
	Path  string `json:"path"`
}

// Gets the npcs from a JSON and returns a slice
func GetNpcsFromJson() ([]Npc, error) {
	npcs, err := loadJSONData[Npc]("../data/npcs.json")
	if err != nil {
		return []Npc{}, fmt.Errorf("error getting npcs from JSON file")
	}

	return npcs, nil
}
