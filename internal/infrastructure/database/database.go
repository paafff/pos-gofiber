package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"pos-gofiber/internal/config"
	"pos-gofiber/internal/domain/models"

	_ "github.com/lib/pq" // Import driver PostgreSQL untuk database/sql
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	// Load database configuration from config.AppConfig
	dbConfig := config.AppConfig.Database

	// Create the database if it doesn't exist
	created, err := CreateDatabaseIfNotExists(dbConfig)
	if err != nil {
		log.Fatal("could not create database: ", err)
	}

	// Create DSN
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		dbConfig.Host, dbConfig.Username, dbConfig.Password, dbConfig.Name, dbConfig.Port, dbConfig.SSLMode, dbConfig.TimeZone)

	// Logika retry untuk koneksi database
	for i := 0; i < 10; i++ {
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		log.Printf("failed to connect to database (attempt %d): %v", i+1, err)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatal("could not connect to the database: ", err)
	}

	// AutoMigrate will create the table based on the User struct
	err = DB.AutoMigrate(&models.User{}, &models.Product{}, &models.OrderedProduct{}, &models.Transaction{})
	if err != nil {
		log.Fatal("failed to migrate database: ", err)
	}

	// Run seeders only if the database was created
	if created {
		SeedUsers(DB, 10)
		SeedProducts(DB, 10)
	}

	log.Println("Database migrated successfully")
}

func CreateDatabaseIfNotExists(dbConfig config.DatabaseConfig) (bool, error) {
	// Create DSN without specifying the database name
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%d sslmode=%s TimeZone=%s",
		dbConfig.Host, dbConfig.Username, dbConfig.Password, dbConfig.Port, dbConfig.SSLMode, dbConfig.TimeZone)

	// Connect to the PostgreSQL server
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return false, err
	}
	defer db.Close()

	// Check if the database exists
	var exists bool
	query := fmt.Sprintf("SELECT EXISTS(SELECT datname FROM pg_catalog.pg_database WHERE datname = '%s')", dbConfig.Name)
	err = db.QueryRow(query).Scan(&exists)
	if err != nil {
		return false, err
	}

	// Create the database if it doesn't exist
	if !exists {
		_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s", dbConfig.Name))
		if err != nil {
			return false, err
		}
		log.Printf("Database %s created successfully", dbConfig.Name)
		return true, nil
	} else {
		log.Printf("Database %s already exists", dbConfig.Name)
	}

	return false, nil
}
