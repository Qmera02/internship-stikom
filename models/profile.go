package models

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	UserID  uint   `json:"user_id"`
	Nama    string `json:"nama"`
	Alamat  string `json:"alamat"`
	Bio     string `json:"bio"`
	Email   string `json:"email"`
	Telepon string `json:"telepon"`
}
