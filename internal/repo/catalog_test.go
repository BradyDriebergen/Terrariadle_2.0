package repo

import (
	"context"
	"terrariadle/internal/db"
	"terrariadle/internal/domain"
	"testing"

	"github.com/google/go-cmp/cmp"
)

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
		t.Errorf("answer mismatch (-want +got):\n%s", diff)
	}
}

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
		t.Errorf("answer mismatch (-want +got):\n%s", diff)
	}
}

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
		t.Errorf("answer mismatch (-want +got):\n%s", diff)
	}
}

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
		t.Errorf("answer mismatch (-want +got):\n%s", diff)
	}
}

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
		t.Errorf("answer mismatch (-want +got):\n%s", diff)
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

func generateCategories() []category {
	return []category{}
}

func generateNpcs() []npc {
	return []npc{}
}

func generateEnemies() []enemy {
	return []enemy{}
}

func generateTriviaQuestions() []triviaQuestion {
	return []triviaQuestion{}
}
