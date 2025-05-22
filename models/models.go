package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

type Vehicle struct {
	gorm.Model
	Name          string  `gorm:"not null"`
	Type          string  `gorm:"not null"`
	Status        string  `gorm:"not null;default:'available'"`
	ImageURL      string  `gorm:"not null"`
	BatteryLevel  int     `gorm:"not null;default:100"`
	Latitude      float64 `gorm:"not null"`
	Longitude     float64 `gorm:"not null"`
	CostPerMinute float64 `gorm:"not null"`
}

type Rental struct {
	gorm.Model
	UserID    uint      `gorm:"not null"`
	VehicleID uint      `gorm:"not null"`
	StartTime time.Time `gorm:"not null"`
	EndTime   time.Time
	Cost      float64
}

type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type VehicleResponse struct {
	ID            uint     `json:"id"`
	Name          string   `json:"name"`
	Type          string   `json:"type"`
	Status        string   `json:"status"`
	Image         string   `json:"image"`
	Battery       int      `json:"battery"`
	Location      Location `json:"location"`
	CostPerMinute float64  `json:"cost_per_minute"`
}
