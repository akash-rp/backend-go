package auth

import (
	"github.com/gin-gonic/gin"
)

func AddAuthRoutes(rg *gin.RouterGroup) {
	rg.POST("/login", loginUser)
	rg.GET("/flow/login", getLoginAuthFlow)
	rg.GET("/flow/logout", getLogoutAuthFlow)
	rg.GET("/logout", logoutUser)
	rg.GET("/getUserDetails", getUserDetails)
	// server := rg.Group("/server")
}
