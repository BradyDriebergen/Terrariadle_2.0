package types

type NPC struct {
	NPC     string   `json:"npc"`
	NPCPath string   `json:"npc-path"`
	Quotes  []string `json:"quotes"`
	Names   []string `json:"names"`
}
