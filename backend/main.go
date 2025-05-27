package main

import (
	"github.com/mjavadtavakoli/fullstack/config"

	"github.com/mjavadtavakoli/fullstack/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.User{})

	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello from backend!"})
	})

	r.Run(":8080")
}
