package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func sendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-type", "application/jason")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s successfull", op),
		"data": data,
	})

}

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-type", "application/jason")
	ctx.JSON(code, gin.H{
		"message": msg,
		"errorCode": code,
	})
}