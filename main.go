package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(AuthMiddleware)
	r.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"test": "success",
		})
	})
	r.Run("127.0.0.1:4200")
}
