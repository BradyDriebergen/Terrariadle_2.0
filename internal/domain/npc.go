package domain

type Npc struct {
	ID      int
	NPC     string
	NPCPath string
	Quotes  []string
	Names   []string
}

type SearchNpcResult struct {
	NpcID int
	Name  string
	Path  string
}
