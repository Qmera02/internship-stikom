package controllers

import (
	"internship-stikom/config"
	"internship-stikom/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateInternship(c *gin.Context) {
	var internship models.Internship
	userID := c.MustGet("user_id").(uint)
	if err := c.ShouldBindJSON(&internship); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	internship.UserID = userID
	if err := config.DB.Create(&internship).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create internship"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "internship created successfully", "internship": internship})
}
func GetInternships(c *gin.Context) {
	UserID := c.MustGet("user_id").(uint)
	var internship []models.Internship
	if err := config.DB.Where("user_id = ?", UserID).Find(&internship).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve internships"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"internships": internship})
}
