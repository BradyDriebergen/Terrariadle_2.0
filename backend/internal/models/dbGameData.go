package models

import (
	"terrariadle-backend/internal/types"
	"time"
)

type DbGameData struct {
	DsData DailySlashData    `bson:"dsData"`
	TcData [4]types.Category `bson:"tcData"`
	GnData GuessNpcData      `bson:"gnData"`
	HmPos  int               `bson:"hmPos"`
	// 7 little words storage
	LastReset time.Time `bson:"lastReset"`
}

type DailySlashData struct {
	CurrentWeaponId  int `bson:"currentWeaponId"`
	PreviousWeaponId int `bson:"previousWeaponId"`
}

type GuessNpcData struct {
	NpcId         int      `bson:"npcId"`
	NpcQuotePos   int      `bson:"npcQuotePos"`
	NpcNamePos    int      `bson:"npcNamePos"`
	OtherNpcNames []string `bson:"otherNpcNames"`
}
