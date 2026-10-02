package testutils

import (
	"math/rand/v2"
	"slices"
	"testing"
)

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

func Shuffle[T any](items []T, seed uint64) []T {
	shuffled := slices.Clone(items)
	r := rand.New(rand.NewPCG(seed, seed))
	r.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})
	return shuffled
}
