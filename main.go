package main

import (
	"os"

	"github.com/b3ns44d/cloud-storage/src/controllers"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {

	gin.ForceConsoleColor()

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.POST("/upload", controllers.Upload)

	router.Run(os.Getenv("PORT"))
}
