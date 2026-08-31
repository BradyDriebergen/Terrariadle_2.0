package repo

import (
	"testing"
)

func TestGetWeapons(t *testing.T) {
	// ctx := context.Background()
	// collection := freshCollection(t)
	// answerRepo := NewCatalogRepo(testMongo, collection, "", "", "", "")

}

func generateWeapons() []weapon {
	return []weapon{
		{
			ID:           10,
			Name:         "Amethyst Staff",
			WeaponType:   "Staff",
			ModeObtained: "Pre-HardMode",
			Info: weaponInfo{
				ImagePath:  "/Amethyst_Staff.png",
				DamageType: "Magic",
				Damage:     15,
				UseTime:    "Very Slow",
				Rarity:     "White",
				Operation:  "Manual",
				Material:   "No",
				Obtained: []string{
					"Crafting",
				},
			},
		},
		{
			ID:           2,
			Name:         "Adamantite Glaive",
			WeaponType:   "Spear",
			ModeObtained: "Hardmode",
			Info: weaponInfo{
				ImagePath:  "/Adamantite_Glaive.png",
				DamageType: "Melee",
				Damage:     49,
				UseTime:    "Fast",
				Rarity:     "Light Red",
				Operation:  "Manual",
				Material:   "No",
				Obtained: []string{
					"Crafting",
				},
			},
		},
		{
			ID:           5,
			Name:         "Aerial Bane",
			WeaponType:   "Bow",
			ModeObtained: "Hardmode",
			Info: weaponInfo{
				ImagePath:  "/Aerial_Bane.png",
				DamageType: "Ranged",
				Damage:     39,
				UseTime:    "Average",
				Rarity:     "Yellow",
				Operation:  "Auto",
				Material:   "No",
				Obtained: []string{
					"Drop",
				},
			},
		},
	}
}
