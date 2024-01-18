package main

import (
	"github.com/gin-gonic/gin"
)

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)

	router.Run("localhost:8080")
}
