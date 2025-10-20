package types

type NPCdata struct {
	ID      int      `json:"id"`
	NPC     string   `json:"npc"`
	NPCPath string   `json:"npc-path"`
	Quote   string   `json:"quote"`
	Names   []string `json:"names"`
}

type NPC struct {
	NPC     string   `json:"npc"`
	NPCPath string   `json:"npc-path"`
	Quotes  []string `json:"quotes"`
	Names   []string `json:"names"`
}
