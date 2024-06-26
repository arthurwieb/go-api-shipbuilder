package database_data

import "goshipbuilder/models"

// Returns the default upgrade outfittings that will be inserted into the database if the table is empty.
func GetUpgradeOutfitting() []*models.UpgradeOutfitting {
	upgradeOutfittingData := []*models.UpgradeOutfitting{
		{
			Name:        "Destroyer",
			Hull:        f(10),
			Sails:       f(20),
			Guns:        f(40),
			CannonAcc:   f(5),
			CannonDmg:   f(10),
			RldSpeed:    f(20),
			Greater:     f(10),
			RepairCd:    f(10),
			GunsRp:      f(50),
			BoardChance: f(25),
		},
		{
			Name:      "Dreadnought",
			Hull:      f(20),
			Sails:     f(40),
			Guns:      f(40),
			CannonDmg: f(5),
			Greater:   f(10),
			RepairCd:  f(10),
			HullRp:    f(50),
			SailRp:    f(50),
			GunsRp:    f(50),
			CrewHits:  f(20),
		},
		{
			Name:      "Explorer",
			Sails:     f(40),
			Speed:     f(5),
			HoldCap:   f(100),
			Lesser:    f(10),
			RepairCd:  f(15),
			Doubloons: f(10),
			HealBonus: f(5),
			Fishing:   f(5),
			Spy:       f(20),
			Wake:      f(10),
		},
		{
			Name:      "Fisherman",
			Hull:      f(20),
			Sails:     f(40),
			Speed:     f(2.5),
			Greater:   f(10),
			RepairCd:  f(5),
			SailRp:    f(50),
			Doubloons: f(5),
			Fishing:   f(15),
			Spy:       f(10),
			Wake:      f(20),
		},
		{
			Name:      "Merchant",
			Hull:      f(10),
			Sails:     f(20),
			Speed:     f(2.5),
			HoldCap:   f(200),
			Lesser:    f(10),
			RepairCd:  f(5),
			Doubloons: f(15),
			CrewHits:  f(15),
			HealBonus: f(15),
			Wake:      f(20),
		},
		{
			Name:      "Privateer",
			Sails:     f(20),
			Guns:      f(20),
			Speed:     f(2.5),
			CannonAcc: f(2.5),
			CannonDmg: f(5),
			Lesser:    f(10),
			Regular:   f(10),
			Greater:   f(10),
			Doubloons: f(5),
			Spy:       f(20),
		},
		{
			Name:        "Raider",
			Hull:        f(20),
			Sails:       f(40),
			Speed:       f(2.5),
			Regular:     f(10),
			BoardChance: f(50),
			Doubloons:   f(10),
			CrewHits:    f(20),
			CrewBreach:  f(20),
			CrewDamage:  f(10),
			HealBonus:   f(5),
		},
		{
			Name:       "Runner",
			Sails:      f(40),
			Speed:      f(7.5),
			Lesser:     f(10),
			RepairCd:   f(10),
			SailRp:     f(50),
			CrewBreach: f(10),
			HealBonus:  f(10),
			Spy:        f(10),
		},
		{
			Name:      "Sentry",
			Hull:      f(10),
			Sails:     f(20),
			Speed:     f(2.5),
			CannonAcc: f(5),
			Lesser:    f(10),
			Regular:   f(10),
			Greater:   f(10),
			Spy:       f(15),
			Wake:      f(15),
		},
		{
			Name:      "Skirmisher",
			CannonAcc: f(2.5),
			RldSpeed:  f(20),
			Lesser:    f(20),
			Regular:   f(20),
			Greater:   f(20),
			CrewHits:  f(15),
			HealBonus: f(15),
			Wake:      f(25),
		},
	}
	return upgradeOutfittingData
}
