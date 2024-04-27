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
	ID        uint `gorm:"primaryKey;autoIncrement"`
	Name      string
	Ships     []Ship      `gorm:"foreignKey:TypeID"`
	ShipStats []ShipStats `gorm:"foreignKey:TypeID"`
}

//* significa que pode ser nulo no GORM
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
