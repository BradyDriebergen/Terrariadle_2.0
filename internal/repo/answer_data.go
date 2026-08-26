package repo

import (
	"context"
	"terrariadle-backend/internal/db"
)

type AnswerRepo interface {
	GetAnswerData(ctx context.Context) (AnswerData, error)
	UpsertAnswerData(ctx context.Context, answerData *AnswerData) error
	GetGuessCounts(ctx context.Context) (PlayerGuessCounts, error)
	UpsertGuessCounts(ctx context.Context, guessCounts *PlayerGuessCounts) error
}

type MongoAnswerRepo struct {
	database             *db.MongoDB
	answerCollection     string
	guessCountCollection string
}

func NewAnswerRepo(db *db.MongoDB, aCollection, gcCollection string) *MongoAnswerRepo {
	return &MongoAnswerRepo{
		database:             db,
		answerCollection:     aCollection,
		guessCountCollection: gcCollection,
	}
}

func (r *MongoAnswerRepo) GetAnswerData(ctx context.Context) (AnswerData, error) {
	answerData, err := db.FindOne[AnswerData](ctx, r.database, r.answerCollection, db.Filter{"_id": 1})
	if err != nil {
		return AnswerData{}, err
	}

	return *answerData, nil
}

func (r *MongoAnswerRepo) UpsertAnswerData(ctx context.Context, answerData *AnswerData) error {
	err := db.Upsert(ctx, r.database, r.answerCollection, db.Filter{"_id": 1}, answerData)
	return err
}

func (r *MongoAnswerRepo) GetGuessCounts(ctx context.Context) (PlayerGuessCounts, error) {
	guessCounts, err := db.FindOne[PlayerGuessCounts](ctx, r.database, r.guessCountCollection, db.Filter{"_id": 1})
	if err != nil {
		return PlayerGuessCounts{}, err
	}

	return *guessCounts, nil
}

func (r *MongoAnswerRepo) UpsertGuessCounts(ctx context.Context, guessCounts *PlayerGuessCounts) error {
	err := db.Upsert(ctx, r.database, r.guessCountCollection, db.Filter{"_id": 1}, guessCounts)
	return err
}
