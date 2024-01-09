package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"isaacszf.antiqbrasblog.com/domain"
	"isaacszf.antiqbrasblog.com/domain/authentication"
	"isaacszf.antiqbrasblog.com/domain/models"
	"isaacszf.antiqbrasblog.com/domain/utils"
)

type PostCreateInput struct {
	HeroImage string `json:"hero_image" binding:"required"`
	Title string `json:"title" binding:"required"`
	Subtitle string `json:"subtitle" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func CreatePost(c *gin.Context) {
	var input PostCreateInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	writerID, err := authentication.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	writerName, err := GetWriterNameByID(writerID.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	post := models.Post{
		HeroImage: input.HeroImage,
		Title: input.Title,
		Subtitle: input.Subtitle,
		Content: input.Content,
		Slug: utils.GenerateSlug(input.Title),
		WriterName: writerName,
		WriterID: writerID,
	}

	res := domain.DB.Create(&post)
	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, post)
}
