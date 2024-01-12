package routes

import (
	"backend/controllers"
	"github.com/gin-gonic/gin"
)

func AddServerRoutes(rg *gin.RouterGroup) {
	rg.GET("/servers", controllers.ServersList)
	rg.GET("/server/:id", controllers.ServerDetails)

	//Fail2ban
	rg.GET("/server/:id/fail2ban", controllers.GetBannedIps)

	//Health
	rg.GET("/server/:id/health", controllers.ServerHealth)
	rg.GET("/server/:id/health/:period", controllers.ServerHealth)

	//Services
	rg.GET("/server/:id/services", controllers.ServicesStatus)
	rg.PATCH("/server/:id/service", controllers.ServiceAction)

	//SSH Users
	rg.GET("/server/:id/ssh", controllers.SshUsers)
	rg.DELETE("/server/:id/ssh", controllers.KillSshUser)

	//SSH keys
	rg.GET("/server/:id/sshKeys", controllers.SshKeys)
	rg.POST("/server/:id/sshKey", controllers.AddSshKey)

	//System Users
	rg.GET("/server/:id/users", controllers.SystemUsers)
	rg.PATCH("/server/:id/users", controllers.ChangeUserPassword)
	rg.DELETE("/server/:id/users", controllers.DeleteSystemUser)

	//UFW
	rg.GET("/server/:id/ufw", controllers.UfwRules)
	rg.POST("/server/:id/ufw", controllers.ManageUfwRule)
	rg.DELETE("/server/:id/ufw", controllers.ManageUfwRule)

	rg.GET("/server/:id/sites", controllers.SitesByServer)
}
