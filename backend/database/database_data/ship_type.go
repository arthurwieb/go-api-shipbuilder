package database_data

import (
	"goshipbuilder/models"
)

func f(f float64) *float64 {
	return &f
}

func i(i uint) *uint {
	return &i
}

func GetShipTypes() []*models.ShipType {
	shipTypesData := []*models.ShipType{
		{
			Name:         "Small",
			MaxCrew:      2,
			Hull:         f(2000),
			Sails:        f(1500),
			Guns:         f(1500),
			Speed:        f(6.25),
			CannonMinDmg: f(40),
			CannonMaxDmg: f(60),
			RldSpeed:     f(10),
			Lesser:       f(120),
			Regular:      f(240),
			Greater:      f(260),
			RepairCd:     f(300),
			HullRp:       f(10),
			SailRp:       f(20),
			GunsRp:       f(20),
			BoardChance:  f(25),
			Wake:         f(100),
		},
		{
			Name:         "Medium",
			MaxCrew:      3,
			Hull:         f(3000),
			Sails:        f(2000),
			Guns:         f(2000),
			Speed:        f(6.556),
			CannonMinDmg: f(40),
			CannonMaxDmg: f(60),
			RldSpeed:     f(15),
			Lesser:       f(120),
			Regular:      f(240),
			Greater:      f(260),
			RepairCd:     f(300),
			HullRp:       f(10),
			SailRp:       f(20),
			GunsRp:       f(20),
			BoardChance:  f(25),
			Wake:         f(125),
		},
		{
			Name:         "Large",
			MaxCrew:      4,
			Hull:         f(4000),
			Sails:        f(3000),
			Guns:         f(3000),
			Speed:        f(5),
			CannonMinDmg: f(40),
			CannonMaxDmg: f(60),
			RldSpeed:     f(10),
			Lesser:       f(120),
			Regular:      f(240),
			Greater:      f(260),
			RepairCd:     f(300),
			HullRp:       f(10),
			SailRp:       f(20),
			GunsRp:       f(20),
			BoardChance:  f(25),
			Wake:         f(150),
		},
		{
			Name:         "Carrack",
			MaxCrew:      5,
			Hull:         f(5000),
			Sails:        f(3500),
			Guns:         f(3500),
			Speed:        f(4.545),
			CannonMinDmg: f(40),
			CannonMaxDmg: f(60),
			RldSpeed:     f(10),
			Lesser:       f(120),
			Regular:      f(240),
			Greater:      f(260),
			RepairCd:     f(300),
			HullRp:       f(10),
			SailRp:       f(20),
			GunsRp:       f(20),
			BoardChance:  f(25),
			Wake:         f(175),
		},
		{
			Name:         "Galleon",
			MaxCrew:      6,
			Hull:         f(6000),
			Sails:        f(4500),
			Guns:         f(4500),
			Speed:        f(4.167),
			CannonMinDmg: f(40),
			CannonMaxDmg: f(60),
			RldSpeed:     f(10),
			Lesser:       f(120),
			Regular:      f(240),
			Greater:      f(260),
			RepairCd:     f(300),
			HullRp:       f(10),
			SailRp:       f(20),
			GunsRp:       f(20),
			BoardChance:  f(25),
			Wake:         f(200),
		},
	}
	return shipTypesData
}
