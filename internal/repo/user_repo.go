package repo

import (
	"context"
	"errors"
	"terrariadle/internal/db"
	"terrariadle/internal/domain"
)

type UserRepo interface {
	GetUser(ctx context.Context, userId string) (domain.User, error)
	UpsertUserData(ctx context.Context, user domain.User) error
	DropAllUserData(ctx context.Context) error
}

type MongoUserRepo struct {
	database       *db.MongoDB
	userCollection string
}

func NewUserRepo(db *db.MongoDB, uCollection string) *MongoUserRepo {
	return &MongoUserRepo{
		database:       db,
		userCollection: uCollection,
	}
}

func (r *MongoUserRepo) GetUser(ctx context.Context, userId string) (domain.User, error) {
	user, err := db.FindOne[userData](ctx, r.database, r.userCollection, db.Filter{"userId": userId})
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			return domain.User{}, ErrNotFound
		}
		return domain.User{}, err
	}
	return toUser(*user), nil
}

func (r *MongoUserRepo) UpsertUserData(ctx context.Context, user domain.User) error {
	err := db.Upsert(ctx, r.database, r.userCollection, db.Filter{"userId": user.UserID}, toUserData(user))
	return err
}

func (r *MongoUserRepo) DropAllUserData(ctx context.Context) error {
	err := db.DeleteAll(ctx, r.database, r.userCollection)
	return err
}
