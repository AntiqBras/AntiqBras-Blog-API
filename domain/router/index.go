package router

import (
	"github.com/gin-gonic/gin"
	"isaacszf.antiqbrasblog.com/domain/authentication"
	"isaacszf.antiqbrasblog.com/domain/controllers"
)

func Load(r *gin.Engine) {
	r.GET("/", landing)
	
	r.GET("/posts", controllers.GetAllPosts)
	r.GET("/posts/:slug", controllers.GetPostBySlug)
	r.POST("/posts", authentication.JWTAuthMiddleware(), controllers.CreatePost)
	r.PUT("/posts/:id", authentication.JWTAuthMiddleware(), controllers.UpdatePost)
	r.DELETE("/posts/:id", authentication.JWTAuthMiddleware(), controllers.DeletePost)

	r.POST("/writers/login", controllers.Login)
	r.POST("/writers/register", authentication.JWTAuthMiddleware(), controllers.Register)
	r.GET("/writers/:username", authentication.JWTAuthMiddleware(), controllers.GetWriterByUsername)
}

func landing(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}
