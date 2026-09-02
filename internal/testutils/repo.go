package testutils

import (
	"context"
	"terrariadle/internal/domain"
)

// struct and methods for mocking answer_repo
type FakeAnswerRepo struct {
	AnswerData      domain.AnswerRefs
	GuessCounts     domain.PlayerGuessCounts
	GetAnswerErr    error
	UpsertAnswerErr error
	GetGuessErr     error
	UpsertGuessErr  error
}

func (f *FakeAnswerRepo) GetAnswerData(ctx context.Context) (domain.AnswerRefs, error) {
	if f.GetAnswerErr != nil {
		return domain.AnswerRefs{}, f.GetAnswerErr
	}
	return f.AnswerData, nil
}

func (f *FakeAnswerRepo) UpsertAnswerData(ctx context.Context, answerData *domain.AnswerRefs) error {
	if f.UpsertAnswerErr != nil {
		return f.UpsertAnswerErr
	}
	f.AnswerData = *answerData
	return nil
}

func (f *FakeAnswerRepo) GetGuessCounts(ctx context.Context) (domain.PlayerGuessCounts, error) {
	if f.GetGuessErr != nil {
		return domain.PlayerGuessCounts{}, f.GetGuessErr
	}
	return f.GuessCounts, nil
}

func (f *FakeAnswerRepo) UpsertGuessCounts(ctx context.Context, guessCounts *domain.PlayerGuessCounts) error {
	if f.UpsertGuessErr != nil {
		return f.UpsertGuessErr
	}
	f.GuessCounts = *guessCounts
	return nil
}
