package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	dbUrl := os.Getenv("DB_URL")

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World! with dbUrl: %s", dbUrl)
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
