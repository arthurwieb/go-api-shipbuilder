package database_data

import (
	"goshipbuilder/models"
)

func GetMyShips() []*models.Ship {
	myShipData := []*models.Ship{
		{
			ShipTypeID:          Large,
			UpgradeOutfittingID: 1,
			Name:                "NewJeans",
			HullBonus:           f(9.5),
			SailsBonus:          f(9.5),
			GunsBonus:           f(9.5),
			SpeedBonus:          f(9.5),
			CannonAccBonus:      f(9.5),
			CannonMinDmgBonus:   f(9.5),
			CannonMaxDmgBonus:   f(9.5),
			RldSpeedBonus:       f(9.5),
			HoldCapBonus:        f(9.5),
			LesserBonus:         f(9.5),
			RegularBonus:        f(9.5),
			GreaterBonus:        f(9.5),
			RepairCdBonus:       f(9.5),
			HullRpBonus:         f(9.5),
			SailRpBonus:         f(9.5),
			GunsRpBonus:         f(9.5),
			BoardChanceBonus:    f(9.5),
			DoubloonsBonus:      f(9.5),
			CrewHitsBonus:       f(9.5),
			CrewBreachBonus:     f(9.5),
			CrewDamageBonus:     f(9.5),
			HealBonusBonus:      f(9.5),
			FishingBonus:        f(9.5),
			SpyBonus:            f(9.5),
			WakeBonus:           f(9.5),
		},
	}
	return myShipData
}
