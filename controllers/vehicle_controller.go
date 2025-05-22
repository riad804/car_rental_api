package controllers

import (
	"time"

	"github.com/car_rental_api/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type VehicleController struct {
	DB *gorm.DB
}

func NewVehicleController(db *gorm.DB) *VehicleController {
	return &VehicleController{DB: db}
}

func (vc *VehicleController) GetVehicleList(c *fiber.Ctx) error {
	var vehicles []models.Vehicle
	if err := vc.DB.Find(&vehicles).Error; err != nil {
		return models.Error(c, fiber.StatusInternalServerError, "Could not fetch vehicles")
	}

	response := make([]models.VehicleResponse, len(vehicles))
	for i, vehicle := range vehicles {
		response[i] = models.VehicleResponse{
			ID:            vehicle.ID,
			Name:          vehicle.Name,
			Type:          vehicle.Type,
			Status:        vehicle.Status,
			Image:         vehicle.ImageURL,
			Battery:       vehicle.BatteryLevel,
			Location:      models.Location{Lat: vehicle.Latitude, Lng: vehicle.Longitude},
			CostPerMinute: vehicle.CostPerMinute,
		}
	}

	return models.Success(c, "success", response)
}

func (vc *VehicleController) GetVehicleDetails(c *fiber.Ctx) error {
	vehicleID := c.Params("id")

	var vehicle models.Vehicle
	if err := vc.DB.First(&vehicle, vehicleID).Error; err != nil {
		return models.Error(c, fiber.StatusNotFound, "Vehicle not found")
	}

	response := models.VehicleResponse{
		ID:            vehicle.ID,
		Name:          vehicle.Name,
		Type:          vehicle.Type,
		Status:        vehicle.Status,
		Image:         vehicle.ImageURL,
		Battery:       vehicle.BatteryLevel,
		Location:      models.Location{Lat: vehicle.Latitude, Lng: vehicle.Longitude},
		CostPerMinute: vehicle.CostPerMinute,
	}

	return models.Success(c, "success", response)
}

func (vc *VehicleController) RentVehicle(c *fiber.Ctx) error {
	vehicleID := c.Params("id")
	userID := c.Locals("userID").(uint)

	var vehicle models.Vehicle
	if err := vc.DB.First(&vehicle, vehicleID).Error; err != nil {
		return models.Error(c, fiber.StatusNotFound, "Vehicle not found")
	}

	if vehicle.Status != "available" {
		return models.Error(c, fiber.StatusBadRequest, "Vehicle is not available for rent")
	}

	// Start rental
	rental := models.Rental{
		UserID:    userID,
		VehicleID: vehicle.ID,
		StartTime: time.Now(),
	}

	if err := vc.DB.Create(&rental).Error; err != nil {
		return models.Error(c, fiber.StatusInternalServerError, "Could not start rental")
	}

	// Update vehicle status
	if err := vc.DB.Model(&vehicle).Update("status", "rented").Error; err != nil {
		return models.Error(c, fiber.StatusInternalServerError, "Could not update vehicle status")
	}

	return models.Success(c, "Rental started successfully", fiber.Map{
		"rental_id": rental.ID,
	})
}
