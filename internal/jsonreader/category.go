package jsonreader

import "fmt"

// Catagory struct that gets pulled from JSON files

type Category struct {
	ID       int      `json:"id"`
	Category string   `json:"category"`
	Options  []string `json:"options"`
}

// Gets the categories from a JSON and returns a slice
func GetCategoriesFromJson() ([]Category, error) {
	categories, err := loadJSONData[Category]("../data/categories.json")
	if err != nil {
		return []Category{}, fmt.Errorf("error getting categories from JSON file")
	}

	return categories, nil
}
