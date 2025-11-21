package jsonreader

import "fmt"

// Weapon struct that gets pulled from JSON files

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

type SearchWeaponResult struct {
	WeaponId int    `json:"weaponId"`
	Name     string `json:"name"`
	Path     string `json:"path"`
}

// Gets the weapons from a JSON and returns a slice
func GetWeaponsFromJSON() ([]Weapon, error) {
	weapons, err := loadJSONData[Weapon]("../data/weapons.json")
	if err != nil {
		return []Weapon{}, fmt.Errorf("error getting weapons from JSON file")
	}

	return weapons, nil
}
