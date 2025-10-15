package types

type WeaponOutput struct {
	PreviousWeaponName string `json:"previousWeaponName"`
	Hint1              string `json:"hint1"`
	Hint2              string `json:"hint2"`
	Hint3              string `json:"hing3"`
}

type WeaponData struct {
	CurrentWeapon  Weapon `json:"currentWeapon"`
	PreviousWeapon Weapon `json:"previousWeapon"`
}

type Weapon struct {
	ID           int        `json:"id"`
	Name         string     `json:"name"`
	WeaponType   string     `json:"weapon-type"`
	ModeObtained string     `json:"mode-obtained"`
	Info         WeaponInfo `json:"info"`
}

type WeaponInfo struct {
	ImagePath  string   `json:"image-path"`
	DamageType string   `json:"damage-type"`
	Damage     int      `json:"damage"`
	UseTime    string   `json:"use-time"`
	Rarity     string   `json:"rarity"`
	Operation  string   `json:"operation"`
	Material   string   `json:"material"`
	Obtained   []string `json:"obtained"`
}
