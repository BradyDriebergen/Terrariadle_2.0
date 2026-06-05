package repo

type weapon struct {
	ID           int        `bson:"id"`
	Name         string     `bson:"name"`
	WeaponType   string     `bson:"weapon-type"`
	ModeObtained string     `bson:"mode-obtained"`
	Info         weaponInfo `bson:"info"`
}

type weaponInfo struct {
	ImagePath  string   `bson:"image-path"`
	DamageType string   `bson:"damage-type"`
	Damage     int      `bson:"damage"`
	UseTime    string   `bson:"use-time"`
	Rarity     string   `bson:"rarity"`
	Operation  string   `bson:"operation"`
	Material   string   `bson:"material"`
	Obtained   []string `bson:"obtained"`
}

type category struct {
	ID       int      `bson:"id"`
	Category string   `bson:"category"`
	Options  []string `bson:"options"`
}

type npc struct {
	ID      int      `bson:"id"`
	NPC     string   `bson:"npc"`
	NpcPath string   `bson:"npc-path"`
	Quotes  []string `bson:"quotes"`
	Names   []string `bson:"names"`
}

type enemy struct {
	ID        int    `bson:"id"`
	Name      string `bson:"name"`
	ImagePath string `bson:"image_path"`
}
