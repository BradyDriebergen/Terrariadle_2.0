package repo

import (
	"context"
	"terrariadle/internal/db"
	"terrariadle/internal/domain"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

// Checks if we can get all weapons from the database
func TestGetWeapons(t *testing.T) {
	ctx := context.Background()
	catalogRepo := mockCatalogRepo(t)

	weapons := []weapon{
		newWeapon(1, "Sword"),
		newWeapon(2, "Bow"),
		newWeapon(3, "Staff"),
	}

	want := []domain.Weapon{
		newDomainWeapon(1, "Sword"),
		newDomainWeapon(2, "Bow"),
		newDomainWeapon(3, "Staff"),
	}

	err := db.InsertMany(ctx, testMongo, catalogRepo.weaponCollection, weapons)
	if err != nil {
		t.Fatalf("insertmany failed: %v", err)
	}

	got, err := catalogRepo.GetWeapons(ctx)
	if err != nil {
		t.Fatalf("getweapons failed: %v", err)
	}

	if len(got) != 3 {
		t.Errorf("Expected a length of 3, got %v", len(got))
	}

	less := func(a, b domain.Weapon) bool { return a.ID < b.ID }

	if diff := cmp.Diff(want, got, cmpopts.SortSlices(less)); diff != "" {
		t.Errorf("weapon mismatch (-want +got):\n%s", diff)
	}
}

// Checks if we can get all categories from the database
func TestGetCategories(t *testing.T) {
	ctx := context.Background()
	catalogRepo := mockCatalogRepo(t)

	categories := []category{
		newCategory(1, "Rarity Colors"),
		newCategory(2, "Enemy Types"),
		newCategory(3, "Npc Names"),
	}

	want := []domain.Category{
		newDomainCategory(1, "Rarity Colors"),
		newDomainCategory(2, "Enemy Types"),
		newDomainCategory(3, "Npc Names"),
	}

	err := db.InsertMany(ctx, testMongo, catalogRepo.categoryCollection, categories)
	if err != nil {
		t.Fatalf("insertmany failed: %v", err)
	}

	got, err := catalogRepo.GetCategories(ctx)
	if err != nil {
		t.Fatalf("getcategories failed: %v", err)
	}

	if len(got) != 3 {
		t.Errorf("Expected a length of 3, got %v", len(got))
	}

	less := func(a, b domain.Category) bool { return a.ID < b.ID }

	if diff := cmp.Diff(want, got, cmpopts.SortSlices(less)); diff != "" {
		t.Errorf("catagories mismatch (-want +got):\n%s", diff)
	}
}

// Checks if we can get all npcs from the database
func TestGetNpcs(t *testing.T) {
	ctx := context.Background()
	catalogRepo := mockCatalogRepo(t)

	npcs := []npc{
		newNpc(1, "Guide"),
		newNpc(2, "Nurse"),
		newNpc(3, "Party Girl"),
	}

	want := []domain.Npc{
		newDomainNpc(1, "Guide"),
		newDomainNpc(2, "Nurse"),
		newDomainNpc(3, "Party Girl"),
	}

	err := db.InsertMany(ctx, testMongo, catalogRepo.npcCollection, npcs)
	if err != nil {
		t.Fatalf("insertmany failed: %v", err)
	}

	got, err := catalogRepo.GetNpcs(ctx)
	if err != nil {
		t.Fatalf("getnpcs failed: %v", err)
	}

	if len(got) != 3 {
		t.Errorf("Expected a length of 3, got %v", len(got))
	}

	less := func(a, b domain.Npc) bool { return a.ID < b.ID }

	if diff := cmp.Diff(want, got, cmpopts.SortSlices(less)); diff != "" {
		t.Errorf("npcs mismatch (-want +got):\n%s", diff)
	}
}

// Checks if we can get all enemies from the database
func TestGetEnemies(t *testing.T) {
	ctx := context.Background()
	catalogRepo := mockCatalogRepo(t)

	enemies := []enemy{
		newEnemy(1, "Zombie"),
		newEnemy(2, "Skeleton"),
		newEnemy(3, "Bat"),
	}

	want := []domain.Enemy{
		newDomainEnemy(1, "Zombie"),
		newDomainEnemy(2, "Skeleton"),
		newDomainEnemy(3, "Bat"),
	}

	err := db.InsertMany(ctx, testMongo, catalogRepo.enemyCollection, enemies)
	if err != nil {
		t.Fatalf("insertmany failed: %v", err)
	}

	got, err := catalogRepo.GetEnemies(ctx)
	if err != nil {
		t.Fatalf("getenemies failed: %v", err)
	}

	if len(got) != 3 {
		t.Errorf("Expected a length of 3, got %v", len(got))
	}

	less := func(a, b domain.Enemy) bool { return a.ID < b.ID }

	if diff := cmp.Diff(want, got, cmpopts.SortSlices(less)); diff != "" {
		t.Errorf("enemies mismatch (-want +got):\n%s", diff)
	}
}

// Checks if we can get all trivia questions from the database
func TestGetTriviaQuestions(t *testing.T) {
	ctx := context.Background()
	catalogRepo := mockCatalogRepo(t)

	questions := []triviaQuestion{
		newTriviaQuestion(1, "Green Sword"),
		newTriviaQuestion(2, "Stylist Name"),
		newTriviaQuestion(3, "Common Block"),
	}

	want := []domain.TriviaQuestion{
		newDomainTriviaQuestion(1, "Green Sword"),
		newDomainTriviaQuestion(2, "Stylist Name"),
		newDomainTriviaQuestion(3, "Common Block"),
	}

	err := db.InsertMany(ctx, testMongo, catalogRepo.triviaCollection, questions)
	if err != nil {
		t.Fatalf("insertmany failed: %v", err)
	}

	got, err := catalogRepo.GetTriviaQuestions(ctx)
	if err != nil {
		t.Fatalf("gettriviaquestions failed: %v", err)
	}

	if len(got) != 3 {
		t.Errorf("Expected a length of 3, got %v", len(got))
	}

	less := func(a, b domain.TriviaQuestion) bool { return a.ID < b.ID }

	if diff := cmp.Diff(want, got, cmpopts.SortSlices(less)); diff != "" {
		t.Errorf("trivia questions mismatch (-want +got):\n%s", diff)
	}
}

// Helper method creating collections for catalog_repo
func mockCatalogRepo(t *testing.T) *MongoCatalogRepo {
	t.Helper()

	weaponCollection := "weapon_" + t.Name()
	categoryCollection := "category_" + t.Name()
	npcCollection := "npc_" + t.Name()
	enemyCollection := "enemy_" + t.Name()
	triviaCollection := "trivia_" + t.Name()

	t.Cleanup(func() {
		_ = db.DeleteAll(context.Background(), testMongo, weaponCollection)
		_ = db.DeleteAll(context.Background(), testMongo, categoryCollection)
		_ = db.DeleteAll(context.Background(), testMongo, npcCollection)
		_ = db.DeleteAll(context.Background(), testMongo, enemyCollection)
		_ = db.DeleteAll(context.Background(), testMongo, triviaCollection)
	})

	return NewCatalogRepo(
		testMongo,
		weaponCollection,
		categoryCollection,
		npcCollection,
		enemyCollection,
		triviaCollection,
	)
}

// generates weapon objects for testing
func newDomainWeapon(id int, name string) domain.Weapon {
	return domain.Weapon{ID: id, Name: name}
}

func newWeapon(id int, name string) weapon {
	return weapon{ID: id, Name: name}
}

// generates category objects for testing
func newDomainCategory(id int, name string) domain.Category {
	return domain.Category{ID: id, Category: name}
}

func newCategory(id int, name string) category {
	return category{ID: id, Category: name}
}

// generates npc objects for testing
func newDomainNpc(id int, name string) domain.Npc {
	return domain.Npc{ID: id, NPC: name}
}

func newNpc(id int, name string) npc {
	return npc{ID: id, NPC: name}
}

// generates enemy objects for testing
func newDomainEnemy(id int, name string) domain.Enemy {
	return domain.Enemy{ID: id, Name: name}
}

func newEnemy(id int, name string) enemy {
	return enemy{ID: id, Name: name}
}

// generates trivia question objects for testing
func newDomainTriviaQuestion(id int, name string) domain.TriviaQuestion {
	return domain.TriviaQuestion{ID: id, Clue: name}
}

func newTriviaQuestion(id int, name string) triviaQuestion {
	return triviaQuestion{ID: id, Clue: name}
}
