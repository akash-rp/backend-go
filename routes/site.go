package routes

import (
	"backend/controllers"
	"github.com/gin-gonic/gin"
)

func AddSitesRoutes(rg *gin.RouterGroup) {
	rg.POST("/:id", controllers.AddSite)
	rg.GET("/:id", controllers.GetSiteById)
	rg.POST("/:id/domain", controllers.AddDomain)
	rg.POST("/:id/domain/primary", controllers.ChangePrimaryDomain)
	rg.DELETE("/:id/domain/:domainId", controllers.DeleteDomain)
	rg.PATCH("/:id/domain", controllers.DomainWildcard)
	rg.GET("/:id/php/ini", controllers.PhpIni)
	rg.PATCH("/:id/php/ini", controllers.UpdatePhpIni)
	rg.GET("/:id/php/settings", controllers.PhpSettings)
	rg.PATCH("/:id/php/Settings", controllers.UpdatePhpSettings)
	rg.GET("/:id/firewall", controllers.FirewallBySiteId)
	rg.POST("/:id/firewall/7g", controllers.UpdateSevenGFirewall)
	rg.POST("/:id/firewall/modsecurity", controllers.UpdateModSecFirewall)
}
