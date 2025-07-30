package controllers

import (
	"internship-stikom/config"
	"internship-stikom/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProfile(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found"})
		return
	}

	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":              user.ID,
		"user_name":       user.Name,
		"email":           user.Email,
		"profile_name":    user.Profile.Nama,
		"profile_address": user.Profile.Alamat,
	})
}

func CreateProfile(c *gin.Context) {
	var profile models.Profile
	userID := c.MustGet("user_id").(uint)

	if err := c.ShouldBindJSON(&profile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	profile.UserID = userID

	if err := config.DB.Create(&profile).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile berhasil dibuat", "data": profile})
}
