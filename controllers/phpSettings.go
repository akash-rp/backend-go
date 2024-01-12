package controllers

import (
	"backend/db"
	"backend/models"
	"backend/site"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PhpSettings(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	siteId := ctx.Param("id")

	tx, err := db.DbConn.Begin(ctx)
	if err != nil {
		fmt.Print(err.Error())
		ctx.AbortWithStatus(400)
		return
	}
	defer tx.Commit(ctx)

	site, err := site.GetSiteDetails(siteId, userId.(string), tx)
	if err != nil {
		fmt.Print(err.Error())
		ctx.AbortWithStatus(400)
		return
	}

	resp, err := http.Get(fmt.Sprintf("http://%s:8081/getPHPsettings/%s", site.IP, site.Name))
	if err != nil {
		fmt.Print(err)
	}
	defer resp.Body.Close()

	var phpSettings models.PhpSettings
	json.NewDecoder(resp.Body).Decode(&phpSettings)

	ctx.JSON(200, phpSettings)
}
