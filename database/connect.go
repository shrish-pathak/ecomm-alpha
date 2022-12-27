package database

import (
	"ecomm-alpha/config"
	"ecomm-alpha/models"
	"fmt"
	"log"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectDB connect to db
func ConnectDB() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	log.Println("Connection Opened to Database")
	DB.AutoMigrate(
		// &models.Product{},
		// &models.Buyer{},
		&models.Seller{},
		&models.Store{},
	// &models.Addresses{},
	)
	log.Println("Database Migrated")
}
