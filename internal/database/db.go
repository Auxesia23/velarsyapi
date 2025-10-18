package database

import (
	"fmt"
	"log"
	"os"

	"github.com/Auxesia23/velarsyapi/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgres() (*gorm.DB, error) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		return nil, fmt.Errorf("DATABASE_URL environment variable is not set")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.OurWork{}, &models.Service{}, &models.Project{}, &models.Image{}, &models.User{})
	if err != nil {
		return nil, err
	}
	
	seedDatabase(db)

	log.Println("Database initialized")

	return db, nil
}
