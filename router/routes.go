package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine){
	v1 := router.Group( "/api/v1") 
	{
		// Show Opening
		v1.GET("/opening", func(ctx *gin.Context){
			ctx.JSON(http.StatusOK, gin.H{
				"message":"GET Opening",
			})
		})
		// Post Opening
		v1.POST("/opening", func(ctx *gin.Context){
			ctx.JSON(http.StatusOK, gin.H{
				"message":"POST Opening",
			})
		})
		// Delete Opening
		v1.DELETE("/opening", func(ctx *gin.Context){
			ctx.JSON(http.StatusOK, gin.H{
				"message":"DELETE Opening",
			})
		})
		// UpDating Opening
		v1.PUT("/opening", func(ctx *gin.Context){
			ctx.JSON(http.StatusOK, gin.H{
			"message":"PUT Opening",
			})
		})
		// Show Openings
		v1.GET("/openings", func(ctx *gin.Context){
			ctx.JSON(http.StatusOK, gin.H{
			"message":"GET Openings",
			})
		})
	}
}
