package controllers

import (
	"backend/db"
	"backend/models"
	"backend/site"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdatePhpSettings(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	siteId := ctx.Param("id")

	updatePhpSettings := new(models.UpdatePhpSettings)
	if err := ctx.ShouldBind(&updatePhpSettings); err != nil {
		fmt.Print(err.Error())
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	tx, err := db.DbConn.Begin(ctx)
	if err != nil {
		fmt.Print(err.Error())
		ctx.AbortWithStatus(400)
		return
	}
	defer tx.Commit(ctx)

	wp, err := site.GetSiteDetails(siteId, userId.(string), tx)
	if err != nil {
		fmt.Print(err.Error())
		ctx.AbortWithStatus(400)
		return
	}

	updatePhpSettings.User = wp.User

	updatePhpSettingsJSON, _ := json.Marshal(updatePhpSettings)
	resp, err := http.Post(fmt.Sprintf("http://%s:8081/updatePHPsettings/%s", wp.IP, wp.Name), "application/json", bytes.NewReader(updatePhpSettingsJSON))
	if err != nil {
		fmt.Print(err)
	}
	defer resp.Body.Close()

	var phpSettings models.PhpSettings
	json.NewDecoder(resp.Body).Decode(&phpSettings)

	ctx.JSON(200, phpSettings)
}
