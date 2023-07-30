package main

import (
	"backend/auth"
	"backend/server"

	"github.com/gin-gonic/gin"
)

func addRoutes(r *gin.Engine) {
	rgAuth := r.Group("/auth")
	auth.AddAuthRoutes(rgAuth)
	rg := r.Group("/", AuthMiddleware)
	server.AddServerRoutes(rg)
}
