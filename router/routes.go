package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine){
	v1 := router.Group( "/api/v1")
	{
		v1.GET("/opening", func(ctx *gin.Context){
			ctx.JSON(http.StatusOK, gin.H{
				"message":"GET Opening",
			})
		})
	}
}