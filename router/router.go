package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
	// Inizializza il Router
	router := gin.Default()
	// Definisce il percorso
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// Inizializza API
	router.Run(":8080")
}
