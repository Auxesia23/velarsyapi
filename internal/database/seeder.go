package database

import (
	"log"
	"os"

	"github.com/Auxesia23/velarsyapi/internal/models"
	"github.com/Auxesia23/velarsyapi/internal/utils"
	"gorm.io/gorm"
)

func seedDatabase(db *gorm.DB) {
	seedUser(db)
}

func seedUser(db *gorm.DB) {
	var user models.User
	userUsernameSeed := os.Getenv("USER_USERNAME_SEED")
	userPasswordSeed := os.Getenv("USER_PASSWORD_SEED")

	result := db.Where("username = ?", userUsernameSeed).First(&user)
	if result.Error != nil && result.Error == gorm.ErrRecordNotFound {
		hashedPassword, _ := utils.HashPassword(userPasswordSeed)
		newUser := models.User{
			Username: "velarsy",
			Password: hashedPassword,
		}

		if err := db.Create(&newUser).Error; err != nil {
			log.Fatal("Failed to create user")
		}
		log.Println("User created succesfuly")
	} else if result.Error != nil {
		log.Fatal("Failed checking user data", result.Error)
	} else {
		log.Println("User already exist, Seeder skipped")
	}
}
