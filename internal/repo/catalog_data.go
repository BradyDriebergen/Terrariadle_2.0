package repo

import (
	"context"
	"terrariadle-backend/internal/db"
	"terrariadle-backend/internal/domain"
)

type CatalogRepo struct {
	database *db.MongoDB
}

func NewCatalogRepo(db *db.MongoDB) *CatalogRepo {
	return &CatalogRepo{
		database: db,
	}
}

func (r *CatalogRepo) GetWeapons(ctx context.Context) ([]domain.Weapon, error) {
	weaponCatalog, err := db.GetAll[weapon](ctx, r.database, "daily_slash_weapons")
	if err != nil {
		return []domain.Weapon{}, err
	}

	return toDomainWeapons(weaponCatalog), nil
}

func (r *CatalogRepo) GetCategories(ctx context.Context) ([]domain.Category, error) {
	categoryCatalog, err := db.GetAll[category](ctx, r.database, "connections_categories")
	if err != nil {
		return []domain.Category{}, err
	}

	return toDomainCategories(categoryCatalog), nil
}

func (r *CatalogRepo) GetNpcs(ctx context.Context) ([]domain.Npc, error) {
	npcCatalog, err := db.GetAll[npc](ctx, r.database, "guess_the_npc_npcs")
	if err != nil {
		return []domain.Npc{}, err
	}

	return toDomainNpcs(npcCatalog), nil
}

func (r *CatalogRepo) GetEnemies(ctx context.Context) ([]domain.Enemy, error) {
	enemyCatalog, err := db.GetAll[enemy](ctx, r.database, "hangman_enemies")
	if err != nil {
		return []domain.Enemy{}, err
	}

	return toDomainEnemies(enemyCatalog), nil
}
