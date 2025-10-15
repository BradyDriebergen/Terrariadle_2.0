package services

import (
	"fmt"
	"terrariadle-backend/internal/types"
	"terrariadle-backend/internal/utils"
)

func GetDailySlashPuzzleData() (types.WeaponOutput, error) {
	gameData, ok := utils.GetMemData()
	if !ok {
		return types.WeaponOutput{}, fmt.Errorf("no game data cached")
	}

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
