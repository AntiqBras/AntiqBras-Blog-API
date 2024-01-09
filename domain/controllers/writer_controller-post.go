package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"isaacszf.antiqbrasblog.com/domain"
	"isaacszf.antiqbrasblog.com/domain/authentication"
	"isaacszf.antiqbrasblog.com/domain/models"
)

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Author string `json:"author" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	w := models.Writer{}

	w.Username = input.Username
	w.Author = input.Author
	w.Password = input.Password

	err := domain.DB.Create(&w).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "writer registered with success"})
}

func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	w := models.Writer{}

	w.Username = input.Username
	w.Password = input.Password

	token, err := authentication.LoginCheck(w.Username, w.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is invalid"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
