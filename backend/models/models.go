package models

type UpgradeType struct {
	ID   uint `gorm:"primaryKey;autoIncrement"`
	Name string
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

type Ship struct {
	ID                  uint  `gorm:"primaryKey;autoIncrement"`
	ShipTypeID          uint  `gorm:"foreignKey:ID"`
	UpgradeOutfittingID *uint `gorm:"foreignKey:ID"`
	UpgradeSuppliesID   *uint `gorm:"foreignKey:ID"`
	UpgradeSpecialtyID  *uint `gorm:"foreignKey:ID"`
	Name                string
	HullBonus           *float64
	SailsBonus          *float64
	GunsBonus           *float64
	SpeedBonus          *float64
	CannonAccBonus      *float64
	CannonMinDmgBonus   *float64
	CannonMaxDmgBonus   *float64
	RldSpeedBonus       *float64
	HoldCapBonus        *float64
	LesserBonus         *float64
	RegularBonus        *float64
	GreaterBonus        *float64
	RepairCdBonus       *float64
	HullRpBonus         *float64
	SailRpBonus         *float64
	GunsRpBonus         *float64
	BoardChanceBonus    *float64
	DoubloonsBonus      *float64
	CrewHitsBonus       *float64
	CrewBreachBonus     *float64
	CrewDamageBonus     *float64
	HealBonusBonus      *float64
	FishingBonus        *float64
	SpyBonus            *float64
	WakeBonus           *float64
	ShipType            ShipType
	UpgradeOutfitting   UpgradeOutfitting
	UpgradeSpecialty    UpgradeSpecialty
	UpgradeSupplies     UpgradeSupplies
}

type UpgradeOutfitting struct {
	ID          uint `gorm:"primaryKey;autoIncrement"`
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
}
type UpgradeSpecialty struct {
	ID          uint `gorm:"primaryKey;autoIncrement"`
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
}
type UpgradeSupplies struct {
	ID          uint `gorm:"primaryKey;autoIncrement"`
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
}
