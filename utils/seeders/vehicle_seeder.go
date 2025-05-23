package seeders

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/car_rental_api/models"
	"gorm.io/gorm"
)

// VehicleSeeder populates the database with sample vehicle data
func VehicleSeeder(db *gorm.DB) error {
	// Only seed if the vehicles table is empty
	var count int64
	db.Model(&models.Vehicle{}).Count(&count)
	if count > 0 {
		log.Println("Vehicles already exist, skipping seeding")
		return nil
	}

	// Sample vehicle data
	vehicles := []models.Vehicle{
		{
			Name:          "Toyota Axio",
			Type:          "car",
			Status:        "available",
			ImageURL:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTTVxLpt10MZtDLSubc9IURN74AvD2To_xi1g&s",
			BatteryLevel:  rand.Intn(50) + 50,
			Latitude:      23.7635,
			Longitude:     90.3892,
			CostPerMinute: 0.30,
		},
		{
			Name:          "Honda Civic",
			Type:          "car",
			Status:        "unavailable",
			ImageURL:      "https://automobiles.honda.com/-/media/Honda-Automobiles/Vehicles/2025/civic-sedan/non-VLP/10-Family/MY25_Civic_Family_Card_Jelly_Hybrid_2x.jpg?sc_lang=en",
			BatteryLevel:  rand.Intn(50) + 50,
			Latitude:      23.7731,
			Longitude:     90.4127,
			CostPerMinute: 0.27,
		},
		{
			Name:          "Nissan Sunny",
			Type:          "car",
			Status:        "available",
			ImageURL:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTLwyzOoq-qQWIjI7HDBcPXMcc4MV-mr5-Onw&s",
			BatteryLevel:  rand.Intn(50) + 50,
			Latitude:      23.7564,
			Longitude:     90.3978,
			CostPerMinute: 0.22,
		},
		{
			Name:          "Toyota Allion",
			Type:          "car",
			Status:        "available",
			ImageURL:      "https://catalogphoto.goo-net.com/carphoto/10101058_200809r.jpg",
			BatteryLevel:  rand.Intn(50) + 50,
			Latitude:      23.7435,
			Longitude:     90.4201,
			CostPerMinute: 0.31,
		},
		{
			Name:          "Hyundai Elantra",
			Type:          "car",
			Status:        "unavailable",
			ImageURL:      "https://di-uploads-pod9.dealerinspire.com/lynneshyundai/uploads/2024/10/2025-Hyundai-Elantra.png",
			BatteryLevel:  rand.Intn(50) + 50,
			Latitude:      23.7742,
			Longitude:     90.3654,
			CostPerMinute: 0.29,
		},
		{
			Name:          "Suzuki Swift",
			Type:          "car",
			Status:        "available",
			ImageURL:      "https://media.umbraco.io/suzuki-gb/5ccm0d5d/10816_suzuki_swift_508_r1.jpg",
			BatteryLevel:  rand.Intn(50) + 50,
			Latitude:      23.7158,
			Longitude:     90.4242,
			CostPerMinute: 0.23,
		},
		{
			Name:          "Toyota Premio",
			Type:          "car",
			Status:        "unavailable",
			ImageURL:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQg7IPaY6Ablr5Erw1mx9Of1RdhC3majtUP3A&s",
			BatteryLevel:  rand.Intn(50) + 50,
			Latitude:      23.7801,
			Longitude:     90.3765,
			CostPerMinute: 0.35,
		},
		{
			Name:          "Mitsubishi Lancer",
			Type:          "car",
			Status:        "available",
			ImageURL:      "https://www.motoroids.com/wp-content/uploads/2014/12/Mitsubishi-Lancer-Evolution-X-final-concept-e1419932519650-1200x752.jpg",
			BatteryLevel:  rand.Intn(50) + 50,
			Latitude:      23.7495,
			Longitude:     90.4012,
			CostPerMinute: 0.25,
		},
		{
			Name:          "BMW 3 Series",
			Type:          "car",
			Status:        "unavailable",
			ImageURL:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQTuTHAvOZhe9ytOzVfOF5WI8dArFXlGEfDLA&s",
			BatteryLevel:  rand.Intn(50) + 50,
			Latitude:      23.7639,
			Longitude:     90.4406,
			CostPerMinute: 0.50,
		},
		{
			Name:          "Mercedes C-Class",
			Type:          "car",
			Status:        "available",
			ImageURL:      "https://imgd.aeplcdn.com/664x374/n/cw/ec/178535/c-class-exterior-right-front-three-quarter.jpeg?isig=0&q=80",
			BatteryLevel:  rand.Intn(50) + 50,
			Latitude:      23.7951,
			Longitude:     90.4287,
			CostPerMinute: 0.45,
		},
		{
			Name:          "Ford Mustang",
			Type:          "car",
			Status:        "unavailable",
			ImageURL:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTWB1YNdqEj9jTTGOwuHFpY8GUCZHcvnq1RgA&s",
			BatteryLevel:  rand.Intn(50) + 50,
			Latitude:      23.7682,
			Longitude:     90.3850,
			CostPerMinute: 0.48,
		},
		{
			Name:          "Mazda Axela",
			Type:          "car",
			Status:        "available",
			ImageURL:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSURn2VjjzEKpY5ak-z90GwzuWokU2HJQHctw&s",
			BatteryLevel:  rand.Intn(50) + 50,
			Latitude:      23.7272,
			Longitude:     90.4383,
			CostPerMinute: 0.26,
		},
		{
			Name:          "Audi A4",
			Type:          "car",
			Status:        "available",
			ImageURL:      "https://imgd.aeplcdn.com/664x374/n/cw/ec/51909/a4-exterior-right-front-three-quarter-2.jpeg?q=80",
			BatteryLevel:  rand.Intn(50) + 50,
			Latitude:      23.7900,
			Longitude:     90.3702,
			CostPerMinute: 0.49,
		},
		{
			Name:          "Chevrolet Cruze",
			Type:          "car",
			Status:        "unavailable",
			ImageURL:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQdsWjpUNPcKZ7ipemeV8XVyQrVezzpyElTVw&s",
			BatteryLevel:  rand.Intn(50) + 50,
			Latitude:      23.7555,
			Longitude:     90.3593,
			CostPerMinute: 0.33,
		},
		{
			Name:          "Subaru Impreza",
			Type:          "car",
			Status:        "available",
			ImageURL:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQyRnThVP17nwtmJnz-DTklNnh51VThWfln3Q&s",
			BatteryLevel:  rand.Intn(50) + 50,
			Latitude:      23.7432,
			Longitude:     90.4451,
			CostPerMinute: 0.32,
		},
		{
			Name:          "Volkswagen Passat",
			Type:          "car",
			Status:        "unavailable",
			ImageURL:      "https://imgd.aeplcdn.com/664x374/cw/ec/22548/Volkswagen-Passat-Headlamps-135233.jpg?wm=0&q=80",
			BatteryLevel:  rand.Intn(50) + 50,
			Latitude:      23.7689,
			Longitude:     90.4130,
			CostPerMinute: 0.37,
		},
		{
			Name:          "Tesla Model 3",
			Type:          "car",
			Status:        "available",
			ImageURL:      "https://preview.thenewsmarket.com/Previews/NCAP/StillAssets/1920x1080/684698_v3.jpg",
			BatteryLevel:  rand.Intn(50) + 50,
			Latitude:      23.7711,
			Longitude:     90.3905,
			CostPerMinute: 0.50,
		},
	}

	// Seed the database
	result := db.Create(&vehicles)
	if result.Error != nil {
		return fmt.Errorf("failed to seed vehicles: %w", result.Error)
	}

	log.Printf("Seeded %d vehicles successfully\n", len(vehicles))
	return nil
}
