package models

import (
	"time"

	"gorm.io/gorm"
)

type Internship struct {
	gorm.Model
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Company     string    `json:"company"`
	Location    string    `json:"location"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	UserID      uint      `json:"user_id"`
}
