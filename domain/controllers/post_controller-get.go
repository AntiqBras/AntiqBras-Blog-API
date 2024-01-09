package controllers

import (
	"errors"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"isaacszf.antiqbrasblog.com/domain"
	"isaacszf.antiqbrasblog.com/domain/models"
)

func GetAllPosts(c *gin.Context) {
	var pageSize = 5
	var page = 1
	var total int64

	p, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err == nil {
		page = p
	}

	pageSz, err := strconv.Atoi(c.DefaultQuery("pageSize", "5"))
	if err == nil {
		pageSize = pageSz
	}

	if page < 1 {
		page = 1
	}

	if pageSize < 1 {
		pageSize = 5
	}

	offset := (page - 1) * pageSize

	var posts []models.Post
	domain.DB.Model(&models.Post{}).Count(&total)
	domain.DB.
		Limit(pageSize).
		Offset(offset).
		Find(&posts).
		Order("created_at DESC").
		Find(&posts)

	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))

	c.JSON(http.StatusOK, gin.H{
		"data":       posts,
		"total":      total,
		"page":       page,
		"totalPages": totalPages,
		"pageSize":   pageSize,
})
}

func GetPostBySlug(c *gin.Context) {
	var post models.Post

	slug := c.Param("slug")
	res := domain.DB.First(&post, "slug = ?", slug); if res != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "post not found"})
			return
		}
	}

	c.JSON(http.StatusOK, post)
}
