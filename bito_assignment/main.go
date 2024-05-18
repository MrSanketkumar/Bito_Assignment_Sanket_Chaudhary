package main

import (
	"bito_assignment/config"
	"bito_assignment/handlers"
	"bito_assignment/models"
	"bito_assignment/repositories"
	"bito_assignment/routes"
	"bito_assignment/services"
)

func main() {
	// Initialize configuration
	cfg := config.NewConfig()
	cfg.DB.AutoMigrate(&models.Subscription{})

	// Set up repository, service, and handler
	subscriptionRepo := repositories.NewSubscriptionRepository(cfg.DB)
	subscriptionService := services.NewSubscriptionService(subscriptionRepo)
	subscriptionHandler := handlers.NewSubscriptionHandler(subscriptionService)

	// Set up router and start the server
	router := routes.SetupRouter(subscriptionHandler)
	router.Run(":8080")
}
