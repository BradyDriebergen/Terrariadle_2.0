package repo

import (
	"context"
	"terrariadle-backend/internal/db"
)

type AnswerRepo interface {
	GetAnswerData(ctx context.Context) (AnswerData, error)
	UpsertAnswerData(ctx context.Context, answerData *AnswerData) error
}

type MongoAnswerRepo struct {
	database *db.MongoDB
}

func NewAnswerRepo(db *db.MongoDB) *MongoAnswerRepo {
	return &MongoAnswerRepo{
		database: db,
	}
}

func (r *MongoAnswerRepo) GetAnswerData(ctx context.Context) (AnswerData, error) {
	answerData, err := db.FindOne[AnswerData](ctx, r.database, "daily_solutions", db.Filter{"_id": 1})
	if err != nil {
		return AnswerData{}, err
	}

	return *answerData, nil
}

func (r *MongoAnswerRepo) UpsertAnswerData(ctx context.Context, answerData *AnswerData) error {
	err := db.Upsert(ctx, r.database, "daily_solutions", db.Filter{"_id": 1}, answerData)
	return err
}
