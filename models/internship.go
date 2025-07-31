package models

import (
	"gorm.io/gorm"
)

type Internship struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Company     string `json:"company"`
	Location    string `json:"location"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	UserID      uint   `json:"user_id"`
}
