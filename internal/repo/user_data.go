package repo

import (
	"context"
	"errors"
	"terrariadle-backend/internal/db"
	"terrariadle-backend/internal/domain"
)

type UserRepo struct {
	database *db.MongoDB
}

func NewUserRepo(db *db.MongoDB) *UserRepo {
	return &UserRepo{
		database: db,
	}
}

// Tries to get user from db. If user doesn't exist, create a new one.
func (r *UserRepo) GetUser(ctx context.Context, userId string) (domain.User, error) {
	user, err := db.FindOne[userData](ctx, r.database, "user_data", db.Filter{"userId": userId})
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			return domain.User{}, ErrNotFound
		}
		return domain.User{}, err
	}

	return user.toDomain(), nil
}

func (r *UserRepo) UpsertUserData(ctx context.Context, user *domain.User) error {
	err := db.Upsert(ctx, r.database, "user_data", db.Filter{"userId": user.UserID}, fromDomain(*user))
	return err
}

func (r *UserRepo) DropAllUserData(ctx context.Context) error {
	err := db.Drop(ctx, r.database, "user_data")
	return err
}
