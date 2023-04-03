package database

import (
	"fmt"
	"log"
	"os"

	"github.com/RaghibA/pokemon-tcg-go/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var PokeDb DbInstance

func Connect() {
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=%v sslmode=disable TimeZone=UTC",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect DB: %v", err)
		os.Exit(2)
	}

	log.Println("DB connected")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running DB migrations...")
	db.AutoMigrate(&models.User{}, &models.Card{})

	PokeDb = DbInstance{
		Db: db,
	}
}
