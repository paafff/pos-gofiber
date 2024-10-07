// package main

// import (
//     "log"
//     "pos-gofiber/internal/auth"
//     "pos-gofiber/internal/domain/services"
//     "pos-gofiber/internal/handlers"
//     "pos-gofiber/internal/infrastructure/database"
//     "pos-gofiber/internal/repositories"
// )

// func initializeDependencies() (*handlers.UserHandler, *auth.AuthHandler, *handlers.ProductHandler, *handlers.TransactionHandler, *handlers.OrderedProductHandler) {
//     // Inisialisasi repository
//     userRepository := repositories.NewUserRepository(database.DB)
//     productRepository := repositories.NewProductRepository(database.DB)
//     transactionRepository := repositories.NewTransactionRepository(database.DB)
//     orderedProductRepository := repositories.NewOrderedProductRepository(database.DB)
//     log.Println("Repositories initialized")

//     // Inisialisasi service
//     userService := services.NewUserService(userRepository)
//     authService := auth.NewAuthService(userRepository)
//     productService := services.NewProductService(productRepository)
//     transactionService := services.NewTransactionService(transactionRepository)
//     orderedProductService := services.NewOrderedProductService(orderedProductRepository)
//     log.Println("Services initialized")

//     // Inisialisasi handler
//     userHandler := handlers.NewUserHandler(userService)
//     authHandler := auth.NewAuthHandler(authService)
//     productHandler := handlers.NewProductHandler(productService)
//     transactionHandler := handlers.NewTransactionHandler(transactionService)
//     orderedProductHandler := handlers.NewOrderedProductHandler(orderedProductService)
//     log.Println("Handlers initialized")

//     return userHandler, authHandler, productHandler, transactionHandler, orderedProductHandler
// }