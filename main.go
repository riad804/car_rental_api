package main

import (
	"log"

	"github.com/car_rental_api/config"
	"github.com/car_rental_api/routes"
	"github.com/car_rental_api/utils/seeders"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Initialize database
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Create Fiber app
	app := fiber.New()

	// Middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))
	app.Use(logger.New())

	// Setup routes
	routes.SetupRoutes(app, db)

	// Seed database with sample data
	if err := seeders.VehicleSeeder(db); err != nil {
		log.Fatalf("Vehicle seeding failed: %v", err)
	}
	log.Println("Vehicle seeding completed successfully")

	// Start server
	log.Fatal(app.Listen(":8080"))
}
