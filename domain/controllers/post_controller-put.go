package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"isaacszf.antiqbrasblog.com/domain"
	"isaacszf.antiqbrasblog.com/domain/models"
	"isaacszf.antiqbrasblog.com/domain/utils"
)

type UpdatePostInput struct {
	HeroImage string `json:"hero_image"`
	Title string `json:"title"`
	Subtitle string `json:"subtitle"`
	Content string `json:"content"`
}

func UpdatePost(c *gin.Context) {
	var input UpdatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var post models.Post

	id := c.Param("id")
	err := domain.DB.Where("id = ?", id).First(&post).Error; if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} 

	if input.HeroImage != "" {
		post.HeroImage = input.HeroImage
	}
	if input.Title != "" {
		post.Title = input.Title
		post.Slug = utils.GenerateSlug(post.Title)
	}
	if input.Subtitle != "" {
		post.Subtitle = input.Subtitle
	}
	if input.Content != "" {
		post.Content = input.Content
	}

	domain.DB.Save(&post)

	c.JSON(http.StatusOK, post)
}
