package repo

import (
	"context"
	"errors"
	"terrariadle-backend/internal/db"
	"terrariadle-backend/internal/domain"
)

type UserRepo interface {
	GetUser(ctx context.Context, userId string) (domain.User, error)
	UpsertUserData(ctx context.Context, user domain.User) error
	DropAllUserData(ctx context.Context) error
}

type MongoUserRepo struct {
	database *db.MongoDB
}

func NewUserRepo(db *db.MongoDB) *MongoUserRepo {
	return &MongoUserRepo{
		database: db,
	}
}

func (r *MongoUserRepo) GetUser(ctx context.Context, userId string) (domain.User, error) {
	user, err := db.FindOne[userData](ctx, r.database, "user_data", db.Filter{"userId": userId})
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			return domain.User{}, ErrNotFound
		}
		return domain.User{}, err
	}
	return user.toDomain(), nil
}

func (r *MongoUserRepo) UpsertUserData(ctx context.Context, user domain.User) error {
	err := db.Upsert(ctx, r.database, "user_data", db.Filter{"userId": user.UserID}, fromDomain(user))
	return err
}

func (r *MongoUserRepo) DropAllUserData(ctx context.Context) error {
	err := db.Drop(ctx, r.database, "user_data")
	return err
}
