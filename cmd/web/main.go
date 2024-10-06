package main

import (
	"fmt"
	"log"
	"os"
	"pos-gofiber/internal/auth"
	"pos-gofiber/internal/config"
	"pos-gofiber/internal/domain/services"
	"pos-gofiber/internal/handlers"
	"pos-gofiber/internal/infrastructure/database"
	"pos-gofiber/internal/infrastructure/router"
	"pos-gofiber/internal/repositories"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load configuration
	config.LoadConfig()
	log.Println("Configuration loaded")

	// Inisialisasi database
	database.InitDatabase()
	log.Println("Database initialized")

	// Handle command line arguments
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "seed":
			database.SeedUsers(database.DB, 10)
			database.SeedProducts(database.DB, 10)
			fmt.Println("Seeded 10 data")
			return
		case "reset-db":
			database.ResetDatabase(database.DB)
			fmt.Println("Database reset.")
			return
		}
	}

	// Inisialisasi repository
	userRepository := repositories.NewUserRepository(database.DB)
	productRepository := repositories.NewProductRepository(database.DB)
	transactionRepository := repositories.NewTransactionRepository(database.DB)
	orderedProductRepository := repositories.NewOrderedProductRepository(database.DB)
	log.Println("initialized")

	// Inisialisasi service
	userService := services.NewUserService(userRepository)
	authService := auth.NewAuthService(userRepository)
	productService := services.NewProductService(productRepository)
	transactionService := services.NewTransactionService(transactionRepository)
	orderedProductService := services.NewOrderedProductService(orderedProductRepository)
	log.Println("service initialized")

	// Inisialisasi handler
	userHandler := handlers.NewUserHandler(userService)
	authHandler := auth.NewAuthHandler(authService)
	productHandler := handlers.NewProductHandler(productService)
	transactionHandler := handlers.NewTransactionHandler(transactionService)
	orderedProductHandler := handlers.NewOrderedProductHandler(orderedProductService)
	log.Println("handler initialized")

	// // Inisialisasi service
	// userService := services.NewUserService(userRepository)
	// log.Println("User service initialized")

	// // Inisialisasi handler
	// userHandler := handlers.NewUserHandler(userService)
	// log.Println("User handler initialized")

	// Inisialisasi Fiber
	app := fiber.New()
	log.Println("Fiber app initialized")

	// Setup routes
	router.SetupRoutes(app, userHandler, authHandler, productHandler, transactionHandler, orderedProductHandler)
	log.Println("Routes set up")

	// Jalankan server
	log.Println("Server running on port 5000")
	if err := app.Listen(":5000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
