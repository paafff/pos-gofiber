package database

import (
	"log"
	"math/rand"
	"time"

	"pos-gofiber/internal/domain/models"

	"github.com/bxcodec/faker/v3"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB, count int) {
	password := "123456"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Could not hash password: %v", err)
	}

	for i := 0; i < count; i++ {
		user := models.User{
			Name:     faker.Name(),
			Email:    faker.Email(),
			Password: string(hashedPassword),
		}
		if err := db.Create(&user).Error; err != nil {
			log.Fatalf("Could not seed user: %v", err)
		}
	}
	log.Printf("Seeded %d users", count)
}

func SeedProducts(db *gorm.DB, count int) {
	for i := 0; i < count; i++ {
		product := models.Product{
			Name:        faker.Word(),
			CreatedByID: uint(rand.Intn(10) + 1), // Random user ID between 1 and 10
			ImageUrl:    faker.URL(),
			Stock:       rand.Intn(100),                // Random stock between 0 and 100
			Sold:        rand.Intn(50),                 // Random sold between 0 and 50
			Price:       float64(rand.Intn(900) + 100), // Random price between 100 and 1000
			Description: faker.Sentence(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
		if err := db.Create(&product).Error; err != nil {
			log.Fatalf("Could not seed product: %v", err)
		}
	}
	log.Printf("Seeded %d products", count)
}

// Fungsi untuk mereset database
func ResetDatabase(db *gorm.DB) {
	err := db.Migrator().DropTable(&models.User{})
	if err != nil {
		log.Fatalf("Could not drop table: %v", err)
	}
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Could not migrate table: %v", err)
	}
	log.Println("Database reset.")
}
