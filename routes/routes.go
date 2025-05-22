package routes

import (
	"github.com/car_rental_api/controllers"
	"github.com/car_rental_api/middlewares"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	// Initialize controllers
	authController := controllers.NewAuthController(db)
	vehicleController := controllers.NewVehicleController(db)
	userController := controllers.NewUserController(db)

	// Auth routes
	auth := app.Group("/auth")
	auth.Post("/login", authController.Login)
	auth.Post("/register", authController.Register)

	// Vehicle routes
	api := app.Group("/api", middlewares.AuthMiddleware)
	api.Get("/vehicles", vehicleController.GetVehicleList)
	api.Get("/vehicles/:id", vehicleController.GetVehicleDetails)

	api.Post("/vehicles/:id/rent", vehicleController.RentVehicle)
	api.Get("/profile/:id", userController.GetProfile)
}
