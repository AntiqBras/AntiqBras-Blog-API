package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"isaacszf.antiqbrasblog.com/domain"
	"isaacszf.antiqbrasblog.com/domain/models"
)

func DeletePost(c *gin.Context) {
	var post models.Post

	id := c.Param("id")
	err := domain.DB.Where("id = ?", id).First(&post).Error; if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} 

	domain.DB.Delete(&post)
	c.JSON(http.StatusOK, 
		gin.H{"status": fmt.Sprintf("post with id=%v was deleted with success", id)})
}
