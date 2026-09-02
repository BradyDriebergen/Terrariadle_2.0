package repo

import (
	"context"
	"terrariadle/internal/db"
	"terrariadle/internal/domain"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Checks if we can get all weapons from the database
func TestGetWeapons(t *testing.T) {
	ctx := context.Background()
	catalogRepo := mockCatalogRepo(t)

	weapons := generateWeapons()
	want := make([]domain.Weapon, len(weapons))
	for i := range weapons {
		want[i] = toWeapon(weapons[i])
	}

	err := db.InsertMany(ctx, testMongo, catalogRepo.weaponCollection, weapons)
	if err != nil {
		t.Fatalf("insertmany failed: %v", err)
	}

	got, err := catalogRepo.GetWeapons(ctx)
	if err != nil {
		t.Fatalf("getweapons failed: %v", err)
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("weapon mismatch (-want +got):\n%s", diff)
	}
}

// Checks if we can get all categories from the database
func TestGetCategories(t *testing.T) {
	ctx := context.Background()
	catalogRepo := mockCatalogRepo(t)

	categories := generateCategories()
	want := make([]domain.Category, len(categories))
	for i := range categories {
		want[i] = toCategory(categories[i])
	}

	err := db.InsertMany(ctx, testMongo, catalogRepo.categoryCollection, categories)
	if err != nil {
		t.Fatalf("insertmany failed: %v", err)
	}

	got, err := catalogRepo.GetCategories(ctx)
	if err != nil {
		t.Fatalf("getcategories failed: %v", err)
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("category mismatch (-want +got):\n%s", diff)
	}
}

// Checks if we can get all npcs from the database
func TestGetNpcs(t *testing.T) {
	ctx := context.Background()
	catalogRepo := mockCatalogRepo(t)

	npcs := generateNpcs()
	want := make([]domain.Npc, len(npcs))
	for i := range npcs {
		want[i] = toNpc(npcs[i])
	}

	err := db.InsertMany(ctx, testMongo, catalogRepo.npcCollection, npcs)
	if err != nil {
		t.Fatalf("insertmany failed: %v", err)
	}

	got, err := catalogRepo.GetNpcs(ctx)
	if err != nil {
		t.Fatalf("getnpcs failed: %v", err)
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("npc mismatch (-want +got):\n%s", diff)
	}
}

// Checks if we can get all enemies from the database
func TestGetEnemies(t *testing.T) {
	ctx := context.Background()
	catalogRepo := mockCatalogRepo(t)

	enemies := generateEnemies()
	want := make([]domain.Enemy, len(enemies))
	for i := range enemies {
		want[i] = toEnemy(enemies[i])
	}

	err := db.InsertMany(ctx, testMongo, catalogRepo.enemyCollection, enemies)
	if err != nil {
		t.Fatalf("insertmany failed: %v", err)
	}

	got, err := catalogRepo.GetEnemies(ctx)
	if err != nil {
		t.Fatalf("getenemies failed: %v", err)
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("enemy mismatch (-want +got):\n%s", diff)
	}
}

// Checks if we can get all trivia questions from the database
func TestGetTriviaQuestions(t *testing.T) {
	ctx := context.Background()
	catalogRepo := mockCatalogRepo(t)

	questions := generateTriviaQuestions()
	want := make([]domain.TriviaQuestion, len(questions))
	for i := range questions {
		want[i] = toTriviaQuestion(questions[i])
	}

	err := db.InsertMany(ctx, testMongo, catalogRepo.triviaCollection, questions)
	if err != nil {
		t.Fatalf("insertmany failed: %v", err)
	}

	got, err := catalogRepo.GetTriviaQuestions(ctx)
	if err != nil {
		t.Fatalf("gettriviaquestions failed: %v", err)
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("trivia question mismatch (-want +got):\n%s", diff)
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
func generateWeapons() []weapon {
	return []weapon{
		{
			ID:           1,
			Name:         "weapon",
			WeaponType:   "type",
			ModeObtained: "hardmode",
			Info: weaponInfo{
				ImagePath:  ".png",
				DamageType: "type",
				Damage:     15,
				UseTime:    "time",
				Rarity:     "rarity",
				Operation:  "operation",
				Material:   "no",
				Obtained:   []string{"obtained"},
			},
		},
		{
			ID:           2,
			Name:         "weapon",
			WeaponType:   "type",
			ModeObtained: "hardmode",
			Info: weaponInfo{
				ImagePath:  ".png",
				DamageType: "type",
				Damage:     15,
				UseTime:    "time",
				Rarity:     "rarity",
				Operation:  "operation",
				Material:   "no",
				Obtained:   []string{"obtained"},
			},
		},
		{
			ID:           3,
			Name:         "weapon",
			WeaponType:   "type",
			ModeObtained: "hardmode",
			Info: weaponInfo{
				ImagePath:  ".png",
				DamageType: "type",
				Damage:     15,
				UseTime:    "time",
				Rarity:     "rarity",
				Operation:  "operation",
				Material:   "no",
				Obtained:   []string{"obtained"},
			},
		},
	}
}

// generates category objects for testing
func generateCategories() []category {
	return []category{
		{
			ID:       1,
			Category: "cat1",
			Options:  []string{"opt1", "opt2", "opt3", "opt4"},
		},
		{
			ID:       2,
			Category: "cat2",
			Options:  []string{"opt1", "opt2", "opt3", "opt4"},
		},
		{
			ID:       3,
			Category: "cat3",
			Options:  []string{"opt1", "opt2", "opt3", "opt4"},
		},
	}
}

// generates npc objects for testing
func generateNpcs() []npc {
	return []npc{
		{
			ID:      1,
			NPC:     "npc1",
			NpcPath: ".png",
			Quotes:  []string{"quote1", "quote2"},
			Names:   []string{"name1", "name2"},
		},
		{
			ID:      2,
			NPC:     "npc2",
			NpcPath: ".png",
			Quotes:  []string{"quote1", "quote2"},
			Names:   []string{"name1", "name2"},
		},
		{
			ID:      3,
			NPC:     "npc3",
			NpcPath: ".png",
			Quotes:  []string{"quote1", "quote2"},
			Names:   []string{"name1", "name2"},
		},
	}
}

// generates enemy objects for testing
func generateEnemies() []enemy {
	return []enemy{
		{
			ID:        1,
			Name:      "enemy",
			ImagePath: ".png",
		},
		{
			ID:        2,
			Name:      "enemy",
			ImagePath: ".png",
		},
		{
			ID:        3,
			Name:      "enemy",
			ImagePath: ".png",
		},
	}
}

// generates trivia weapon objects for testing
func generateTriviaQuestions() []triviaQuestion {
	return []triviaQuestion{
		{
			ID:         1,
			Answer:     "answer1",
			Clue:       "clue",
			Chunks:     []string{"chunk1", "chunk2"},
			ChunkCount: 2,
		},
		{
			ID:         2,
			Answer:     "answer2",
			Clue:       "clue",
			Chunks:     []string{"chunk1", "chunk2"},
			ChunkCount: 2,
		},
		{
			ID:         3,
			Answer:     "answer3",
			Clue:       "clue",
			Chunks:     []string{"chunk1", "chunk2"},
			ChunkCount: 2,
		},
	}
}
