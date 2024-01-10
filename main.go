package main

import (
	"log"

	"isaacszf.antiqbrasblog.com/domain"
	"isaacszf.antiqbrasblog.com/domain/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env: ", err)
	}

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
}))

	domain.Connect()
	router.Load(r)

	r.Run()
}