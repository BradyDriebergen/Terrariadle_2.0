package jsonreader

import (
	"encoding/json"
	"fmt"
	"os"
)

// Helper method for reading from JSON files
func loadJSONData[T any](path string) ([]T, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var data []T
	if err := json.Unmarshal(bytes, &data); err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, fmt.Errorf("no data in %v", path)
	}

	return data, nil
}
