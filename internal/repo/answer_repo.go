package repo

import (
	"context"
	"terrariadle/internal/db"
	"terrariadle/internal/domain"
)

type AnswerRepo interface {
	GetAnswerData(ctx context.Context) (domain.AnswerRefs, error)
	UpsertAnswerData(ctx context.Context, answerData *domain.AnswerRefs) error
	GetGuessCounts(ctx context.Context) (domain.PlayerGuessCounts, error)
	UpsertGuessCounts(ctx context.Context, guessCounts *domain.PlayerGuessCounts) error
}

type MongoAnswerRepo struct {
	database             *db.MongoDB
	answerCollection     string
	guessCountCollection string
}

func NewAnswerRepo(
	db *db.MongoDB,
	aCollection,
	gcCollection string,
) *MongoAnswerRepo {
	return &MongoAnswerRepo{
		database:             db,
		answerCollection:     aCollection,
		guessCountCollection: gcCollection,
	}
}

func (r *MongoAnswerRepo) GetAnswerData(ctx context.Context) (domain.AnswerRefs, error) {
	ad, err := db.FindOne[answerData](ctx, r.database, r.answerCollection, db.Filter{"_id": 1})
	if err != nil {
		return domain.AnswerRefs{}, err
	}

	return toAnswerRef(*ad), nil
}

func (r *MongoAnswerRepo) UpsertAnswerData(ctx context.Context, da *domain.AnswerRefs) error {
	err := db.Upsert(ctx, r.database, r.answerCollection, db.Filter{"_id": 1}, toAnswerData(*da))
	return err
}

func (r *MongoAnswerRepo) GetGuessCounts(ctx context.Context) (domain.PlayerGuessCounts, error) {
	guessCounts, err := db.FindOne[guessCounts](ctx, r.database, r.guessCountCollection, db.Filter{"_id": 1})
	if err != nil {
		return domain.PlayerGuessCounts{}, err
	}

	return toPlayerGuessCounts(*guessCounts), nil
}

func (r *MongoAnswerRepo) UpsertGuessCounts(ctx context.Context, gc *domain.PlayerGuessCounts) error {
	guessCounts := toGuessCounts(*gc)

	err := db.Upsert(ctx, r.database, r.guessCountCollection, db.Filter{"_id": 1}, guessCounts)
	return err
}
