package models

import "gorm.io/gorm"

type Fact struct {
	gorm.Model
	Question string `json:"question" gorm:"text;not null;default:null`
	Answer   string `json:"answer" gorm:"text;not null;default:null`
}

type UpgradeType struct {
	ID           uint `gorm:"primaryKey;autoIncrement"`
	Name         string
	ShipUpgrades []Upgrade `gorm:"foreignKey:TypeID"`
}

type ShipType struct {
	ID           uint `gorm:"primaryKey;autoIncrement"`
	Name         string
	MaxCrew      uint
	Hull         *float64
	Sails        *float64
	Guns         *float64
	Speed        *float64
	CannonAcc    *float64
	CannonMinDmg *float64
	CannonMaxDmg *float64
	RldSpeed     *float64
	HoldCap      *float64
	Lesser       *float64
	Regular      *float64
	Greater      *float64
	RepairCd     *float64
	HullRp       *float64
	SailRp       *float64
	GunsRp       *float64
	BoardChance  *float64
	Doubloons    *float64
	CrewHits     *float64
	CrewBreach   *float64
	CrewDamage   *float64
	HealBonus    *float64
	Fishing      *float64
	Spy          *float64
	Wake         *float64
}

// * significa que pode ser nulo no GORM
type Upgrade struct {
	ID          uint `gorm:"primaryKey;autoIncrement"`
	Name        string
	TypeID      uint
	Type        UpgradeType `gorm:"foreignKey:TypeID"`
	Description string      `gorm:"type:text"`
	Hull        *float64
	Sails       *float64
	Guns        *float64
	Speed       *float64
	CannonAcc   *float64
	CannonDmg   *float64
	RldSpeed    *float64
	HoldCap     *float64
	Lesser      *float64
	Regular     *float64
	Greater     *float64
	RepairCd    *float64
	HullRp      *float64
	SailRp      *float64
	GunsRp      *float64
	BoardChance *float64
	Doubloons   *float64
	CrewHits    *float64
	CrewBreach  *float64
	CrewDamage  *float64
	HealBonus   *float64
	Fishing     *float64
	Spy         *float64
	Wake        *float64
}

type Ship struct {
	ID                uint `gorm:"primaryKey;autoIncrement"`
	TypeID            uint
	Type              ShipType `gorm:"foreignKey:TypeID"`
	Hull              *float64
	Sails             *float64
	Guns              *float64
	Speed             *float64
	CannonAcc         *float64
	CannonMinDmg      *float64
	CannonMaxDmg      *float64
	RldSpeed          *float64
	HoldCap           *float64
	Lesser            *float64
	Regular           *float64
	Greater           *float64
	RepairCd          *float64
	HullRp            *float64
	SailRp            *float64
	GunsRp            *float64
	BoardChance       *float64
	Doubloons         *float64
	CrewHits          *float64
	CrewBreach        *float64
	CrewDamage        *float64
	HealBonus         *float64
	Fishing           *float64
	Spy               *float64
	Wake              *float64
	HullBonus         *float64
	SailsBonus        *float64
	GunsBonus         *float64
	SpeedBonus        *float64
	CannonAccBonus    *float64
	CannonMinDmgBonus *float64
	CannonMaxDmgBonus *float64
	RldSpeedBonus     *float64
	HoldCapBonus      *float64
	LesserBonus       *float64
	RegularBonus      *float64
	GreaterBonus      *float64
	RepairCdBonus     *float64
	HullRpBonus       *float64
	SailRpBonus       *float64
	GunsRpBonus       *float64
	BoardChanceBonus  *float64
	DoubloonsBonus    *float64
	CrewHitsBonus     *float64
	CrewBreachBonus   *float64
	CrewDamageBonus   *float64
	HealBonusBonus    *float64
	FishingBonus      *float64
	SpyBonus          *float64
	WakeBonus         *float64
}

type ShipStats struct {
	ID           uint `gorm:"primaryKey;autoIncrement"`
	TypeID       uint
	Type         ShipType `gorm:"foreignKey:TypeID"`
	Hull         *float64
	Sails        *float64
	Guns         *float64
	Speed        *float64
	CannonAcc    *float64
	CannonMinDmg *float64
	CannonMaxDmg *float64
	RldSpeed     *float64
	HoldCap      *float64
	Lesser       *float64
	Regular      *float64
	Greater      *float64
	RepairCd     *float64
	HullRp       *float64
	SailRp       *float64
	GunsRp       *float64
	BoardChance  *float64
	Doubloons    *float64
	CrewHits     *float64
	CrewBreach   *float64
	CrewDamage   *float64
	HealBonus    *float64
	Fishing      *float64
	Spy          *float64
	Wake         *float64
}

type UpgradeOutfitting struct {
	ID uint `gorm:"primaryKey;autoIncrement"`
	// TypeID      uint
	Name        string
	Hull        *float64
	Sails       *float64
	Guns        *float64
	CannonAcc   *float64
	CannonDmg   *float64
	RldSpeed    *float64
	HoldCap     *float64
	Lesser      *float64
	Greater     *float64
	RepairCd    *float64
	BoardChance *float64
	GunsRp      *float64
	HullRp      *float64
	SailRp      *float64
	CrewHits    *float64
	Speed       *float64
	Doubloons   *float64
	HealBonus   *float64
	Fishing     *float64
	Spy         *float64
	Wake        *float64
	Regular     *float64
	CrewBreach  *float64
	CrewDamage  *float64
	HolderCap   *float64
}

// im creating the struct for this type of data
// const outfittings = [
//   {
//     name: 'Destroyer',
//     typeId: UpgradeTypeID.Outfitting,
//     hull: 10,
//     sails: 20,
//     guns: 40,
//     cannonAcc: 5,
//     cannonDmg: 10,
//     rldSpeed: 20,
//     greater: 10,
//     repairCd: 10,
//     gunsRp: 50,
//     boardChance: 25,
//   },
//   {
//     name: 'Dreadnought',
//     typeId: UpgradeTypeID.Outfitting,
//     hull: 20,
//     sails: 40,
//     guns: 40,
//     cannonDmg: 5,
//     greater: 10,
//     repairCd: 10,
//     hullRp: 50,
//     sailRp: 50,
//     gunsRp: 50,
//     crewHits: 20,
//   },
//   {
//     name: 'Explorer',
//     typeId: UpgradeTypeID.Outfitting,
//     sails: 40,
//     speed: 5,
//     holdCap: 100,
//     lesser: 10,
//     repairCd: 15,
//     doubloons: 10,
//     healBonus: 5,
//     fishing: 5,
//     spy: 20,
//     wake: 10,
//   },
//   {
//     name: 'Fisherman',
//     typeId: UpgradeTypeID.Outfitting,
//     hull: 20,
//     sails: 40,
//     speed: 2.5,
//     greater: 10,
//     repairCd: 5,
//     sailRp: 50,
//     doubloons: 5,
//     fishing: 15,
//     spy: 10,
//     wake: 20,
//   },
//   {
//     name: 'Merchant',
//     typeId: UpgradeTypeID.Outfitting,
//     hull: 10,
//     sails: 20,
//     speed: 2.5,
//     holdCap: 200,
//     lesser: 10,
//     repairCd: 5,
//     doubloons: 15,
//     crewHits: 15,
//     healBonus: 15,
//     wake: 20,
//   },
//   {
//     name: 'Privateer',
//     typeId: UpgradeTypeID.Outfitting,
//     sails: 20,
//     guns: 20,
//     speed: 2.5,
//     cannonAcc: 2.5,
//     cannonDmg: 5,
//     lesser: 10,
//     regular: 10,
//     greater: 10,
//     doubloons: 5,
//     spy: 20,
//   },
//   {
//     name: 'Raider',
//     typeId: UpgradeTypeID.Outfitting,
//     hull: 20,
//     sails: 40,
//     speed: 2.5,
//     regular: 10,
//     boardChance: 50,
//     doubloons: 10,
//     crewHits: 20,
//     crewBreach: 20,
//     crewDamage: 10,
//     healBonus: 5,
//   },
//   {
//     name: 'Runner',
//     typeId: UpgradeTypeID.Outfitting,
//     sails: 40,
//     speed: 7.5,
//     lesser: 10,
//     repairCd: 10,
//     sailRp: 50,
//     crewBreach: 10,
//     healBonus: 10,
//     spy: 10,
//   },
//   {
//     name: 'Sentry',
//     typeId: UpgradeTypeID.Outfitting,
//     hull: 10,
//     sails: 20,
//     guns: 20,
//     speed: 2.5,
//     cannonAcc: 5,
//     lesser: 10,
//     regular: 10,
//     greater: 10,
//     spy: 15,
//     wake: 15,
//   },
//   {
//     name: 'Skirmisher',
//     typeId: UpgradeTypeID.Outfitting,
//     cannonAcc: 2.5,
//     rldSpeed: 20,
//     lesser: 20,
//     regular: 20,
//     greater: 20,
//     crewHits: 15,
//     healBonus: 15,
//     wake: 25,
//   },
// ];
