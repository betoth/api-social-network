package main

import (
	_ "api-social-network/docs" // This line is necessary for go-swagger to find your docs
	"api-social-network/internal/domain/service"
	"api-social-network/internal/infrastructure/config"
	"api-social-network/internal/infrastructure/db/postgres"
	"api-social-network/internal/infrastructure/http/controllers"
	"api-social-network/internal/infrastructure/http/router"
	"fmt"
	"log"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	// Load configuration from .env file
	cfg := config.LoadConfig()

	// Initialize database connection
	db, err := postgres.NewDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize repositories and services
	userRepository := &postgres.UserRepositoryPostgres{DB: db}
	userService := service.UserService{UserRepository: userRepository}

	// Initialize controllers
	userController := &controllers.UserController{UserService: userService}
	authController := &controllers.AuthController{UserService: userService}

	// Initialize router
	r := router.NewRouter(userController, authController)

	// Swagger
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Start the server
	fmt.Println("API listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
