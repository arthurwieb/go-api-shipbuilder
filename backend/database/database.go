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

type Dbinstance struct {
	Db *gorm.DB
}

func Seed(db *gorm.DB) {
	isTableEmpty(db, models.UpgradeType{}, database_data.GetUpgradeTypes())
	isTableEmpty(db, models.ShipType{}, database_data.GetShipTypes())
}

func isTableEmpty(db *gorm.DB, model interface{}, data interface{}) {
	var count int64
	db.Model(model).Count(&count)
	if count == 0 {
		db.Create(data)
	}
}

var DB Dbinstance

func ConnectDb() {
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=America/Sao_Paulo",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("running migrations")
	db.AutoMigrate(&models.Fact{})
	db.AutoMigrate(&models.Upgrade{})
	db.AutoMigrate(&models.UpgradeType{})
	db.AutoMigrate(&models.ShipType{})
	//db.AutoMigrate(&models.Ship{})
	//db.AutoMigrate(&models.ShipStats{})
	Seed(db)

	DB = Dbinstance{
		Db: db,
	}
}
