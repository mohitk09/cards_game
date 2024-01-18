package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/health_check", func(c *gin.Context) {
		c.String(http.StatusOK, "App is reachable")
	})

	router.Run("localhost:8080")
}
