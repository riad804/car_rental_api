package controllers

import (
	"github.com/car_rental_api/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{DB: db}
}

func (uc *UserController) GetProfile(c *fiber.Ctx) error {
	userID := c.Params("id")

	var user models.User
	if err := uc.DB.First(&user, userID).Error; err != nil {
		return models.Error(c, fiber.StatusNotFound, "User not found")
	}

	var rentals []models.Rental
	if err := uc.DB.Where("user_id = ?", userID).Find(&rentals).Error; err != nil {
		return models.Error(c, fiber.StatusInternalServerError, "Could not fetch rentals")
	}

	return models.Success(c, "success", fiber.Map{
		"id":          user.ID,
		"name":        user.Name,
		"email":       user.Email,
		"total_trips": len(rentals),
	})
}
