package jobs

import (
	"context"
	"math/rand"
	"terrariadle-backend/internal/domain"
	"terrariadle-backend/internal/store"
	"terrariadle-backend/internal/utils"
	"time"
)

type PuzzleRefreshJob struct {
	answerStore  *store.AnswerStore
	catalogStore *store.CatalogStore
	rng          *rand.Rand
}

func NewPuzzleRefresh(as *store.AnswerStore, cs *store.CatalogStore) *PuzzleRefreshJob {
	return &PuzzleRefreshJob{
		answerStore:  as,
		catalogStore: cs,
		rng:          rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (j *PuzzleRefreshJob) Start(ctx context.Context) {
	for {
		waitDur := utils.TimeUntilNextMidnightFromNow()

		// Used for quick testing/development
		// waitDur := utils.NextShortTime()

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
		GuessCounts:   domain.PlayerGuessCounts{},
		ResetTime:     now,
		NextResetTime: utils.NextMidnight(now),
	})
}
