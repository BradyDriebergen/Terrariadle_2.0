package domain

type Npc struct {
	ID      int
	NPC     string
	NPCPath string
	Quotes  []string
	Names   []string
}

type SearchNpcResult struct {
	NpcId int
	Name  string
	Path  string
}
