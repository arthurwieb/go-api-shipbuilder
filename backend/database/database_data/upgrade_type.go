package database_data

import (
	"goshipbuilder/models"
)

func GetUpgradeTypes() []*models.UpgradeType {
	upgradeTypesData := []*models.UpgradeType{
		{Name: "Lesser"},
		{Name: "Regular"},
		{Name: "Greater"},
		{Name: "Outfitting"},
		{Name: "Specialty Item"},
		{Name: "Crew Supplies"},
	}
	return upgradeTypesData
}
