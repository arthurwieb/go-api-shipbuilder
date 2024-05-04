package database

import (
	"fmt"
	"log"
	"os"

	"goshipbuilder/database/database_data"
	"goshipbuilder/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Seed(db *gorm.DB) {
	isTableEmpty(db, models.UpgradeType{}, database_data.GetUpgradeTypes())
	isTableEmpty(db, models.ShipType{}, database_data.GetShipTypes())
	isTableEmpty(db, models.UpgradeOutfitting{}, database_data.GetUpgradeOutfitting())
	isTableEmpty(db, models.Ship{}, database_data.GetMyShips())
}

func isTableEmpty(db *gorm.DB, model interface{}, data interface{}) {
	var count int64
	db.Model(model).Count(&count)
	if count == 0 {
		db.Create(data)
	}
}

var DB *gorm.DB

func ConnectDb() {
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=America/Sao_Paulo",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("connected")
	DB.Logger = logger.Default.LogMode(logger.Error)

	log.Println("running migrations")
	DB.AutoMigrate(&models.UpgradeType{})
	DB.AutoMigrate(&models.ShipType{})
	DB.AutoMigrate(&models.UpgradeOutfitting{})
	DB.AutoMigrate(&models.Ship{})
	Seed(DB)
}
