package repo

import (
	"context"
	"strings"
	"terrariadle/internal/db"
	"terrariadle/internal/domain"
)

type AnswerRepo interface {
	GetAnswerData(ctx context.Context) (domain.DailyAnswers, error)
	UpsertAnswerData(ctx context.Context, answerData *domain.DailyAnswers) error
	GetGuessCounts(ctx context.Context) (domain.PlayerGuessCounts, error)
	UpsertGuessCounts(ctx context.Context, guessCounts *domain.PlayerGuessCounts) error
}

type MongoAnswerRepo struct {
	database             *db.MongoDB
	answerCollection     string
	guessCountCollection string
	weaponCollection     string
	categoryCollection   string
	npcCollection        string
	enemyCollection      string
	triviaCollection     string
}

func NewAnswerRepo(
	db *db.MongoDB,
	aCollection,
	gcCollection,
	wCollection,
	cCollection,
	nCollection,
	eCollection,
	tCollection string,
) *MongoAnswerRepo {
	return &MongoAnswerRepo{
		database:             db,
		answerCollection:     aCollection,
		guessCountCollection: gcCollection,
		weaponCollection:     wCollection,
		categoryCollection:   cCollection,
		npcCollection:        nCollection,
		enemyCollection:      eCollection,
		triviaCollection:     tCollection,
	}
}

func (r *MongoAnswerRepo) GetAnswerData(ctx context.Context) (domain.DailyAnswers, error) {
	ad, err := db.FindOne[answerData](ctx, r.database, r.answerCollection, db.Filter{"_id": 1})
	if err != nil {
		return domain.DailyAnswers{}, err
	}

	currentWeapon, err := db.FindOne[weapon](ctx, r.database, r.weaponCollection, db.Filter{"id": ad.DailySlash.CurrentWeaponID})
	if err != nil {
		return domain.DailyAnswers{}, err
	}

	prevWeapon, err := db.FindOne[weapon](ctx, r.database, r.weaponCollection, db.Filter{"id": ad.DailySlash.PrevWeaponID})
	if err != nil {
		return domain.DailyAnswers{}, err
	}

	npc, err := db.FindOne[npc](ctx, r.database, r.npcCollection, db.Filter{"id": ad.GuessTheNpc.NpcID})
	if err != nil {
		return domain.DailyAnswers{}, err
	}

	enemy, err := db.FindOne[enemy](ctx, r.database, r.enemyCollection, db.Filter{"id": ad.Hangman.EnemyID})
	if err != nil {
		return domain.DailyAnswers{}, err
	}
	enemy.Name = strings.ToUpper(enemy.Name)

	triviaQuestions := make([]triviaQuestion, 7)
	for i, q := range ad.TerraTrivia.QuestionIDs {
		question, err := db.FindOne[triviaQuestion](ctx, r.database, r.triviaCollection, db.Filter{"id": q})
		if err != nil {
			return domain.DailyAnswers{}, err
		}

		triviaQuestions[i] = *question
	}

	return domain.DailyAnswers{
		DailySlash: domain.WeaponAnswer{
			CurrentWeapon: toWeapon(*currentWeapon),
			PrevWeapon:    toWeapon(*prevWeapon),
		},
		Connections: domain.ConnectionAnswer{
			CategoryIDs: ad.Connections.CategoryIDs,
			Options:     toCategoryOption(ad.Connections.Options),
		},
		GuessTheNpc: domain.NpcAnswer{
			NpcID:       npc.ID,
			Npc:         npc.NPC,
			Quote:       ad.GuessTheNpc.Quote,
			Name:        ad.GuessTheNpc.Name,
			NameOptions: ad.GuessTheNpc.NameOptions,
		},
		Hangman: domain.HangmanAnswer{
			Enemy: toEnemy(*enemy),
		},
		TerraTrivia: domain.TerraTriviaAnswer{
			Questions: toTriviaQuestions(triviaQuestions),
		},
		ResetTime:     ad.ResetTime,
		NextResetTime: ad.NextResetTime,
	}, nil
}

func (r *MongoAnswerRepo) UpsertAnswerData(ctx context.Context, da *domain.DailyAnswers) error {
	err := db.Upsert(ctx, r.database, r.answerCollection, db.Filter{"_id": 1}, toAnswerData(*da))
	return err
}

func (r *MongoAnswerRepo) GetGuessCounts(ctx context.Context) (domain.PlayerGuessCounts, error) {
	guessCounts, err := db.FindOne[playerGuessCounts](ctx, r.database, r.guessCountCollection, db.Filter{"_id": 1})
	if err != nil {
		return domain.PlayerGuessCounts{}, err
	}

	return domain.PlayerGuessCounts{
		DailySlashCount:  guessCounts.DailySlashCount,
		ConnectionsCount: guessCounts.ConnectionsCount,
		GuessTheNpcCount: guessCounts.GuessTheNpcCount,
		HangmanCount:     guessCounts.HangmanCount,
		TerraTriviaCount: guessCounts.TerraTriviaCount,
	}, nil
}

func (r *MongoAnswerRepo) UpsertGuessCounts(ctx context.Context, gc *domain.PlayerGuessCounts) error {
	guessCounts := playerGuessCounts{
		DailySlashCount:  gc.DailySlashCount,
		ConnectionsCount: gc.ConnectionsCount,
		GuessTheNpcCount: gc.GuessTheNpcCount,
		HangmanCount:     gc.HangmanCount,
		TerraTriviaCount: gc.TerraTriviaCount,
	}

	err := db.Upsert(ctx, r.database, r.guessCountCollection, db.Filter{"_id": 1}, guessCounts)
	return err
}
