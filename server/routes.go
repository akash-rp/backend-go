package server

import (
	"backend/server/ufw"
	"github.com/gin-gonic/gin"
)

func AddServerRoutes(rg *gin.RouterGroup) {
	rg.GET("/servers", getServersList)
	rg.GET("/server/:id", getServerDetails)
	rg.GET("/server/:id/health", getServerHealth)
	rg.POST("/server/:id/health", fetchServerHealthByName)
	rg.GET("/server/:id/service/status", getServicesStatus)
	rg.GET("/server/:id/ufw/rules", ufw.GetUfwRules)
	rg.GET("/server/:id/ssh/users", getSshUsers)
	rg.GET("/server/:id/users", getSystemUsers)
	// server := rg.Group("/server")
}
