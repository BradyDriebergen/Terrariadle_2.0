package repo

import (
	"context"
	"terrariadle/internal/db"
	"terrariadle/internal/domain"
)

type CatalogRepo interface {
	GetWeapons(ctx context.Context) ([]domain.Weapon, error)
	GetCategories(ctx context.Context) ([]domain.Category, error)
	GetNpcs(ctx context.Context) ([]domain.Npc, error)
	GetEnemies(ctx context.Context) ([]domain.Enemy, error)
	GetTriviaQuestions(ctx context.Context) ([]domain.TriviaQuestion, error)
}

type MongoCatalogRepo struct {
	database           *db.MongoDB
	weaponCollection   string
	categoryCollection string
	npcCollection      string
	enemyCollection    string
	triviaCollection   string
}

func NewCatalogRepo(
	db *db.MongoDB,
	wCollection,
	cCollection,
	nCollection,
	eCollection,
	tCollection string,
) *MongoCatalogRepo {
	return &MongoCatalogRepo{
		database:           db,
		weaponCollection:   wCollection,
		categoryCollection: cCollection,
		npcCollection:      nCollection,
		enemyCollection:    eCollection,
		triviaCollection:   tCollection,
	}
}

func (r *MongoCatalogRepo) GetWeapons(ctx context.Context) ([]domain.Weapon, error) {
	weaponCatalog, err := db.GetAll[weapon](ctx, r.database, r.weaponCollection)
	if err != nil {
		return []domain.Weapon{}, err
	}

	return toDomainWeapons(weaponCatalog), nil
}

func (r *MongoCatalogRepo) GetCategories(ctx context.Context) ([]domain.Category, error) {
	categoryCatalog, err := db.GetAll[category](ctx, r.database, r.categoryCollection)
	if err != nil {
		return []domain.Category{}, err
	}

	return toDomainCategories(categoryCatalog), nil
}

func (r *MongoCatalogRepo) GetNpcs(ctx context.Context) ([]domain.Npc, error) {
	npcCatalog, err := db.GetAll[npc](ctx, r.database, r.npcCollection)
	if err != nil {
		return []domain.Npc{}, err
	}

	return toDomainNpcs(npcCatalog), nil
}

func (r *MongoCatalogRepo) GetEnemies(ctx context.Context) ([]domain.Enemy, error) {
	enemyCatalog, err := db.GetAll[enemy](ctx, r.database, r.enemyCollection)
	if err != nil {
		return []domain.Enemy{}, err
	}

	return toDomainEnemies(enemyCatalog), nil
}

func (r *MongoCatalogRepo) GetTriviaQuestions(ctx context.Context) ([]domain.TriviaQuestion, error) {
	triviaCatalog, err := db.GetAll[triviaQuestion](ctx, r.database, r.triviaCollection)
	if err != nil {
		return []domain.TriviaQuestion{}, err
	}

	return toDomainTriviaQuestions(triviaCatalog), nil
}
