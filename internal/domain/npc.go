package domain

type Npc struct {
	ID      int
	NPC     string
	NpcPath string
	Quotes  []string
	Names   []string
}

type SearchNpcResult struct {
	NpcID int
	Name  string
	Path  string
}

type NpcInfo struct {
	Name string
	Path string
}
