package controllers

import (
	"internship-stikom/config"
	"internship-stikom/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProject(c *gin.Context) {
	userID := uint(c.MustGet("UserID").(float64)) // karena MapClaims return float64

	var project models.Project
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	project.UserID = userID
	if err := config.DB.Create(&project).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create project"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Project created successfully",
		"data":    project,
	})
}
func GetProjects(c *gin.Context) {
	userID := uint(c.MustGet("UserID").(float64)) // karena MapClaims return float64
	var projects []models.Project

	if err := config.DB.Where("user_id = ?", userID).Find(&projects).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get projects"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"projects": projects})
}
