package repo

type weapon struct {
	id           int        `bson:"id"`
	name         string     `bson:"name"`
	weaponType   string     `bson:"weapon-type"`
	modeObtained string     `bson:"mode-obtained"`
	info         weaponInfo `bson:"info"`
}

type weaponInfo struct {
	imagePath  string   `bson:"image-path"`
	damageType string   `bson:"damage-type"`
	damage     int      `bson:"damage"`
	useTime    string   `bson:"use-time"`
	rarity     string   `bson:"rarity"`
	operation  string   `bson:"operation"`
	material   string   `bson:"material"`
	obtained   []string `bson:"obtained"`
}

type category struct {
	id       int      `bson:"id"`
	category string   `bson:"category"`
	options  []string `bson:"options"`
}

type npc struct {
	id      int      `bson:"id"`
	npc     string   `bson:"npc"`
	npcPath string   `bson:"npc-path"`
	quotes  []string `bson:"quotes"`
	names   []string `bson:"names"`
}

type enemy struct {
	id        int    `bson:"id"`
	name      string `bson:"name"`
	imagePath string `bson:"image_path"`
}
