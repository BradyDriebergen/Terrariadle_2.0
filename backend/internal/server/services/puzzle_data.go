package services

import (
	"terrariadle-backend/internal/types"
	"terrariadle-backend/internal/utils/cache"
)

func GetDailySlashPuzzleData() (types.WeaponOutput, error) {
	gameData := cache.GetGameData()

	weapons := gameData.DailySlash
	weaponData := types.WeaponOutput{
		PreviousWeaponName: weapons.PreviousWeapon.Name,
		Hint1:              weapons.CurrentWeapon.ModeObtained,
		Hint2:              weapons.CurrentWeapon.WeaponType,
		Hint3:              weapons.CurrentWeapon.Info.ImagePath,
	}

	return weaponData, nil
}

func getConnectionsPuzzleData() {

}

func getGuessTheNpcPuzzleData() {

}

func getHangmanPuzzleData() {

}
