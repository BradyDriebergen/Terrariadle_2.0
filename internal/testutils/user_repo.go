package testutils

import (
	"context"
	"terrariadle/internal/domain"
)

type FakeUserRepo struct {
	User domain.User
}

func (f *FakeUserRepo) GetUser(ctx context.Context, userId string) (domain.User, error) {
	if f.User.UserID == "" {
		return domain.User{}, domain.MongoErrNotFound
	}
	return f.User, nil
}

func (f *FakeUserRepo) UpsertUserData(ctx context.Context, user domain.User) error {
	f.User = user
	return nil
}

func (f *FakeUserRepo) DropAllUserData(ctx context.Context) error {
	f.User = domain.User{}
	return nil
}

func GenerateFakeUserRepo() *FakeUserRepo {
	return &FakeUserRepo{
		User: domain.User{},
	}
}
