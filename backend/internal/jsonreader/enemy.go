package jsonreader

import "fmt"

// Enemy struct that gets pulled from JSON files

type Enemy struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	ImagePath string `json:"image_path"`
}

// Gets the enemies from a JSON and returns a slice
func GetEnemiesFromJson() ([]Enemy, error) {
	enemies, err := loadJSONData[Enemy]("../data/enemies.json")
	if err != nil {
		return []Enemy{}, fmt.Errorf("error getting enemies from JSON file")
	}

	return enemies, nil
}
