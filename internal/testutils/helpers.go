package testutils

import "testing"

func FindFromList[T any](t *testing.T, items []T, id int, getID func(T) int) T {
	t.Helper()
	for _, item := range items {
		if getID(item) == id {
			return item
		}
	}
	t.Fatalf("no item with id %d found in fixture", id)
	var zero T
	return zero
}
