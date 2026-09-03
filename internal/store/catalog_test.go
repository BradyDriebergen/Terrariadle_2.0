package store

import (
	"context"
	"terrariadle/internal/domain"
	"terrariadle/internal/testutils"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

// Tests getting weapons returns weapons
func TestGetWeapons(t *testing.T) {
	ctx := context.Background()

	weapons := testutils.GenerateWeapons()
	npcs := testutils.GenerateNpcs()
	enemies := testutils.GenerateEnemies()
	categories := testutils.GenerateCategories()
	triviaQuestions := testutils.GenerateTriviaQuestions()

	fakeRepo := &testutils.FakeCatalogRepo{
		Weapons:         weapons,
		Npcs:            npcs,
		Enemies:         enemies,
		Categories:      categories,
		TriviaQuestions: triviaQuestions,
	}

	store, err := NewCatalogStore(ctx, fakeRepo)
	if err != nil {
		t.Fatalf("newcatalogstore failed: %v", err)
	}

	got := store.GetWeapons()

	// Conversion from a map -> slice doesn't keep order
	// This helper sort function sorts the slice, ensuring all the data is present
	less := func(a, b domain.Weapon) bool { return a.ID < b.ID }

	if diff := cmp.Diff(weapons, got, cmpopts.SortSlices(less)); diff != "" {
		t.Errorf("catalog store mismatch (-want +got):\n%s", diff)
	}
}

// Tests getting a weapon returns the same weapon
func TestGetWeapon(t *testing.T) {
	ctx := context.Background()

	weapons := testutils.GenerateWeapons()
	npcs := testutils.GenerateNpcs()
	enemies := testutils.GenerateEnemies()
	categories := testutils.GenerateCategories()
	triviaQuestions := testutils.GenerateTriviaQuestions()

	fakeRepo := &testutils.FakeCatalogRepo{
		Weapons:         weapons,
		Npcs:            npcs,
		Enemies:         enemies,
		Categories:      categories,
		TriviaQuestions: triviaQuestions,
	}

	store, err := NewCatalogStore(ctx, fakeRepo)
	if err != nil {
		t.Fatalf("newcatalogstore failed: %v", err)
	}

	got, ok := store.GetWeapon(0)
	if !ok {
		t.Fatalf("Item not found")
	}

	if diff := cmp.Diff(weapons[0], got); diff != "" {
		t.Errorf("catalog store mismatch (-want +got):\n%s", diff)
	}
}

// Tests getting categories returns categories
func TestGetCategories(t *testing.T) {
	ctx := context.Background()

	weapons := testutils.GenerateWeapons()
	npcs := testutils.GenerateNpcs()
	enemies := testutils.GenerateEnemies()
	categories := testutils.GenerateCategories()
	triviaQuestions := testutils.GenerateTriviaQuestions()

	fakeRepo := &testutils.FakeCatalogRepo{
		Weapons:         weapons,
		Npcs:            npcs,
		Enemies:         enemies,
		Categories:      categories,
		TriviaQuestions: triviaQuestions,
	}

	store, err := NewCatalogStore(ctx, fakeRepo)
	if err != nil {
		t.Fatalf("newcatalogstore failed: %v", err)
	}

	got := store.GetCategories()

	// Conversion from a map -> slice doesn't keep order
	// This helper sort function sorts the slice, ensuring all the data is present
	less := func(a, b domain.Category) bool { return a.ID < b.ID }

	if diff := cmp.Diff(categories, got, cmpopts.SortSlices(less)); diff != "" {
		t.Errorf("catalog store mismatch (-want +got):\n%s", diff)
	}
}

// Tests getting a category returns the same category
func TestGetCategory(t *testing.T) {
	ctx := context.Background()

	weapons := testutils.GenerateWeapons()
	npcs := testutils.GenerateNpcs()
	enemies := testutils.GenerateEnemies()
	categories := testutils.GenerateCategories()
	triviaQuestions := testutils.GenerateTriviaQuestions()

	fakeRepo := &testutils.FakeCatalogRepo{
		Weapons:         weapons,
		Npcs:            npcs,
		Enemies:         enemies,
		Categories:      categories,
		TriviaQuestions: triviaQuestions,
	}

	store, err := NewCatalogStore(ctx, fakeRepo)
	if err != nil {
		t.Fatalf("newcatalogstore failed: %v", err)
	}

	got, ok := store.GetCategory(0)
	if !ok {
		t.Fatalf("Item not found")
	}

	if diff := cmp.Diff(categories[0], got); diff != "" {
		t.Errorf("catalog store mismatch (-want +got):\n%s", diff)
	}
}

// Tests getting npcs returns npcs
func TestGetNpcs(t *testing.T) {
	ctx := context.Background()

	weapons := testutils.GenerateWeapons()
	npcs := testutils.GenerateNpcs()
	enemies := testutils.GenerateEnemies()
	categories := testutils.GenerateCategories()
	triviaQuestions := testutils.GenerateTriviaQuestions()

	fakeRepo := &testutils.FakeCatalogRepo{
		Weapons:         weapons,
		Npcs:            npcs,
		Enemies:         enemies,
		Categories:      categories,
		TriviaQuestions: triviaQuestions,
	}

	store, err := NewCatalogStore(ctx, fakeRepo)
	if err != nil {
		t.Fatalf("newcatalogstore failed: %v", err)
	}

	got := store.GetNpcs()

	// Conversion from a map -> slice doesn't keep order
	// This helper sort function sorts the slice, ensuring all the data is present
	less := func(a, b domain.Npc) bool { return a.ID < b.ID }

	if diff := cmp.Diff(npcs, got, cmpopts.SortSlices(less)); diff != "" {
		t.Errorf("catalog store mismatch (-want +got):\n%s", diff)
	}
}

// Tests getting a npc returns the same npc
func TestGetNpc(t *testing.T) {
	ctx := context.Background()

	weapons := testutils.GenerateWeapons()
	npcs := testutils.GenerateNpcs()
	enemies := testutils.GenerateEnemies()
	categories := testutils.GenerateCategories()
	triviaQuestions := testutils.GenerateTriviaQuestions()

	fakeRepo := &testutils.FakeCatalogRepo{
		Weapons:         weapons,
		Npcs:            npcs,
		Enemies:         enemies,
		Categories:      categories,
		TriviaQuestions: triviaQuestions,
	}

	store, err := NewCatalogStore(ctx, fakeRepo)
	if err != nil {
		t.Fatalf("newcatalogstore failed: %v", err)
	}

	got, ok := store.GetNpc(0)
	if !ok {
		t.Fatalf("Item not found")
	}

	if diff := cmp.Diff(npcs[0], got); diff != "" {
		t.Errorf("catalog store mismatch (-want +got):\n%s", diff)
	}
}

// Tests getting enemies returns enemies
func TestGetEnemies(t *testing.T) {
	ctx := context.Background()

	weapons := testutils.GenerateWeapons()
	npcs := testutils.GenerateNpcs()
	enemies := testutils.GenerateEnemies()
	categories := testutils.GenerateCategories()
	triviaQuestions := testutils.GenerateTriviaQuestions()

	fakeRepo := &testutils.FakeCatalogRepo{
		Weapons:         weapons,
		Npcs:            npcs,
		Enemies:         enemies,
		Categories:      categories,
		TriviaQuestions: triviaQuestions,
	}

	store, err := NewCatalogStore(ctx, fakeRepo)
	if err != nil {
		t.Fatalf("newcatalogstore failed: %v", err)
	}

	got := store.GetEnemies()

	// Conversion from a map -> slice doesn't keep order
	// This helper sort function sorts the slice, ensuring all the data is present
	less := func(a, b domain.Enemy) bool { return a.ID < b.ID }

	if diff := cmp.Diff(enemies, got, cmpopts.SortSlices(less)); diff != "" {
		t.Errorf("catalog store mismatch (-want +got):\n%s", diff)
	}
}

// Tests getting a enemy returns the same enemy
func TestGetEnemy(t *testing.T) {
	ctx := context.Background()

	weapons := testutils.GenerateWeapons()
	npcs := testutils.GenerateNpcs()
	enemies := testutils.GenerateEnemies()
	categories := testutils.GenerateCategories()
	triviaQuestions := testutils.GenerateTriviaQuestions()

	fakeRepo := &testutils.FakeCatalogRepo{
		Weapons:         weapons,
		Npcs:            npcs,
		Enemies:         enemies,
		Categories:      categories,
		TriviaQuestions: triviaQuestions,
	}

	store, err := NewCatalogStore(ctx, fakeRepo)
	if err != nil {
		t.Fatalf("newcatalogstore failed: %v", err)
	}

	got, ok := store.GetEnemy(0)
	if !ok {
		t.Fatalf("Item not found")
	}

	if diff := cmp.Diff(enemies[0], got); diff != "" {
		t.Errorf("catalog store mismatch (-want +got):\n%s", diff)
	}
}

// Tests getting trivia questions returns trivia questions
func TestGetTriviaQuestions(t *testing.T) {
	ctx := context.Background()

	weapons := testutils.GenerateWeapons()
	npcs := testutils.GenerateNpcs()
	enemies := testutils.GenerateEnemies()
	categories := testutils.GenerateCategories()
	triviaQuestions := testutils.GenerateTriviaQuestions()

	fakeRepo := &testutils.FakeCatalogRepo{
		Weapons:         weapons,
		Npcs:            npcs,
		Enemies:         enemies,
		Categories:      categories,
		TriviaQuestions: triviaQuestions,
	}

	store, err := NewCatalogStore(ctx, fakeRepo)
	if err != nil {
		t.Fatalf("newcatalogstore failed: %v", err)
	}

	got := store.GetTriviaQuestions()

	// Conversion from a map -> slice doesn't keep order
	// This helper sort function sorts the slice, ensuring all the data is present
	less := func(a, b domain.TriviaQuestion) bool { return a.ID < b.ID }

	if diff := cmp.Diff(triviaQuestions, got, cmpopts.SortSlices(less)); diff != "" {
		t.Errorf("catalog store mismatch (-want +got):\n%s", diff)
	}
}

// Tests getting a trivia question returns the same trivia question
func TestGetTriviaQuestion(t *testing.T) {
	ctx := context.Background()

	weapons := testutils.GenerateWeapons()
	npcs := testutils.GenerateNpcs()
	enemies := testutils.GenerateEnemies()
	categories := testutils.GenerateCategories()
	triviaQuestions := testutils.GenerateTriviaQuestions()

	fakeRepo := &testutils.FakeCatalogRepo{
		Weapons:         weapons,
		Npcs:            npcs,
		Enemies:         enemies,
		Categories:      categories,
		TriviaQuestions: triviaQuestions,
	}

	store, err := NewCatalogStore(ctx, fakeRepo)
	if err != nil {
		t.Fatalf("newcatalogstore failed: %v", err)
	}

	got, ok := store.GetTriviaQuestion(0)
	if !ok {
		t.Fatalf("Item not found")
	}

	if diff := cmp.Diff(triviaQuestions[0], got); diff != "" {
		t.Errorf("catalog store mismatch (-want +got):\n%s", diff)
	}
}
