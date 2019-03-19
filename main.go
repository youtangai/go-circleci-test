package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := setUpRouter()

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func setUpRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	return router
}
