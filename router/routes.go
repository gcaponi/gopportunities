package router

import (
	docs "github.com/gcaponi/gopportunities/docs"
	"github.com/gcaponi/gopportunities/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

)

func initializeRoutes(router *gin.Engine){
	
	// Initialize Handler
	basePath := "/api/v1"
	handler.InitializeHandler()
	docs.SwaggerInfo.BasePath = basePath
	
	v1 := router.Group(basePath) 
	{
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.GET("/openings", handler.ListOpeningsHandler)
	}
	// Initialize Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
