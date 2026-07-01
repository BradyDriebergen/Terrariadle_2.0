package jobs

import (
	"context"
	"math/rand"
	"terrariadle-backend/internal/domain"
	"terrariadle-backend/internal/store"
	"time"
)

type PuzzleRefreshJob struct {
	answerStore     store.AnswerStore
	guessCountStore store.GuessCountsStore
	catalogStore    store.CatalogStore
	userStore       store.UserStore
	rng             *rand.Rand
}

func NewPuzzleRefresh(
	as store.AnswerStore,
	gcs store.GuessCountsStore,
	cs store.CatalogStore,
	us store.UserStore,
) *PuzzleRefreshJob {

	return &PuzzleRefreshJob{
		answerStore:     as,
		catalogStore:    cs,
		guessCountStore: gcs,
		userStore:       us,
		rng:             rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (j *PuzzleRefreshJob) Start(ctx context.Context) {
	// If we missed a reset (e.g. server was down at midnight), refresh immediately.
	if j.answerStore.GetAnswers().NextResetTime.Before(time.Now()) {
		j.refresh(ctx)
	}

	// Testing method for refreshing on startup
	j.refresh((ctx))

	for {
		waitDur := domain.TimeUntilNextMidnightFromNow()

		select {
		case <-ctx.Done():
			return
		case <-time.After(waitDur):
			j.refresh(ctx)
		}
	}
}

func (j *PuzzleRefreshJob) refresh(ctx context.Context) {
	now := time.Now()

	j.answerStore.UpsertAnswers(ctx, domain.DailyAnswers{
		DailySlash:    j.refreshWeapons(),
		Connections:   j.refreshCategories(),
		GuessTheNpc:   j.refreshNpc(),
		Hangman:       j.refreshEnemy(),
		TerraTrivia:   j.refreshTriviaQuestions(),
		ResetTime:     now,
		NextResetTime: domain.NextMidnight(now),
	})

	j.guessCountStore.ResetGuessCounts(ctx)

	j.userStore.DropAllUsers(ctx)
}
