package controllers

import (
	"errors"
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


func GetWriterNameByID(id string) (string, error) {
	var w models.Writer

	if err := domain.DB.Where("id = ?", id).First(&w).Error; err != nil {
		return "", errors.New("failed to get writer, verify if ID is valid")
	}

	return w.Author, nil
}