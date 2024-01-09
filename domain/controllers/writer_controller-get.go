package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"isaacszf.antiqbrasblog.com/domain"
	"isaacszf.antiqbrasblog.com/domain/models"
)

func GetWriterByUsername(c *gin.Context) {
	var w models.Writer

	username := c.Param("username")

	if err := domain.DB.Where("username = ?", username).First(&w).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	result := domain.DB.Preload("Posts").First(&w)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return 
	}

	w.PrepareGive()
	c.JSON(http.StatusOK, w)
}