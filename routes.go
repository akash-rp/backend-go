package main

import (
	"backend/auth"
	"backend/middlewares"
	"backend/routes"
	"github.com/gin-gonic/gin"
)

func addRoutes(r *gin.Engine) {
	rgAuth := r.Group("/auth")
	auth.AddAuthRoutes(rgAuth)
	rg := r.Group("/", AuthMiddleware)
	routes.AddServerRoutes(rg)
	siteGroup := rg.Group("/site", middlewares.SiteVerify)
	routes.AddSitesRoutes(siteGroup)
}
