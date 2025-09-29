package types

type Weapons struct {
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
	ImagePath  string `json:"image-path"`
	DamageType string `json:"damage-type"`
	Damage     int    `json:"damage"`
	UseTime    string `json:"use-time"`
	Rarity     string `json:"rarity"`
	Operation  string `json:"operation"`
	Material   string `json:"material"`
	Obtained   string `json:"obtained"`
}
