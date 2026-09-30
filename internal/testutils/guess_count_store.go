package testutils

import (
	"context"
	"terrariadle/internal/domain"
)

type FakeGuessCountStore struct {
	GuessCounts domain.PlayerGuessCounts
}

func (f *FakeGuessCountStore) GetGuessCounts() domain.PlayerGuessCounts {
	return f.GuessCounts
}

func (f *FakeGuessCountStore) ResetGuessCounts(ctx context.Context) error {
	f.GuessCounts = domain.PlayerGuessCounts{
		DailySlashCount:  0,
		ConnectionsCount: 0,
		GuessTheNpcCount: 0,
		HangmanCount:     0,
		TerraTriviaCount: 0,
	}

	return nil
}

func (f *FakeGuessCountStore) IncrementDailySlashCount(ctx context.Context) (int, error) {
	f.GuessCounts.DailySlashCount++
	return f.GuessCounts.DailySlashCount, nil
}

func (f *FakeGuessCountStore) IncrementConnectionsCount(ctx context.Context) (int, error) {
	f.GuessCounts.ConnectionsCount++
	return f.GuessCounts.ConnectionsCount, nil
}

func (f *FakeGuessCountStore) IncrementGuessTheNpcCount(ctx context.Context) (int, error) {
	f.GuessCounts.GuessTheNpcCount++
	return f.GuessCounts.GuessTheNpcCount, nil
}

func (f *FakeGuessCountStore) IncrementHangmanCount(ctx context.Context) (int, error) {
	f.GuessCounts.HangmanCount++
	return f.GuessCounts.HangmanCount, nil
}

func (f *FakeGuessCountStore) IncrementTerraTriviaCount(ctx context.Context) (int, error) {
	f.GuessCounts.TerraTriviaCount++
	return f.GuessCounts.TerraTriviaCount, nil
}

func GenerateFakeGuessCountStore() *FakeGuessCountStore {
	return &FakeGuessCountStore{
		GuessCounts: domain.PlayerGuessCounts{
			DailySlashCount:  1,
			ConnectionsCount: 2,
			GuessTheNpcCount: 3,
			HangmanCount:     4,
			TerraTriviaCount: 5,
		},
	}
}
