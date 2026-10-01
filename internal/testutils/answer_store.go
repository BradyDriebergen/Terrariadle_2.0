package testutils

import (
	"context"
	"terrariadle/internal/domain"
)

type FakeAnswerStore struct {
	Answers domain.DailyAnswers
}

func (f *FakeAnswerStore) GetAnswers() domain.DailyAnswers {
	return f.Answers
}

func (f *FakeAnswerStore) UpsertAnswers(ctx context.Context, answer domain.DailyAnswers) error {
	f.Answers = answer
	return nil
}

func GenerateFakeAnswerStore() *FakeAnswerStore {
	return &FakeAnswerStore{}
}
