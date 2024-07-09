package router

import (
	"api-social-network/internal/infrastructure/http/controllers"
	"api-social-network/internal/infrastructure/http/middlewares"

	"github.com/gorilla/mux"
)

// NewRouter creates a new mux router with all the necessary routes.
func NewRouter(userController *controllers.UserController, authController *controllers.AuthController) *mux.Router {
	router := mux.NewRouter()

	// Public routes
	router.HandleFunc("/login", authController.Login).Methods("POST")
	router.HandleFunc("/users", userController.CreateUser).Methods("POST")

	// Protected routes
	api := router.PathPrefix("/api").Subrouter()
	api.Use(middlewares.AuthMiddleware)

	api.HandleFunc("/logout", authController.Logout).Methods("POST")
	api.HandleFunc("/users/{id}", userController.GetUser).Methods("GET")
	api.HandleFunc("/users", userController.GetAllUsers).Methods("GET")
	api.HandleFunc("/users/{id}", userController.UpdateUser).Methods("PUT")
	api.HandleFunc("/users/{id}", userController.DeleteUser).Methods("DELETE")

	return router
}
