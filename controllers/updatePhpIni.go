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

func UpdatePhpIni(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	siteId := ctx.Param("id")

	updatePhpIni := new(models.UpdatePhpIni)
	if err := ctx.ShouldBind(&updatePhpIni); err != nil {
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
	updatePhpIniJson, _ := json.Marshal(updatePhpIni)
	resp, err := http.Post(fmt.Sprintf("http://%s:8081/updatePHPini/%s", wp.IP, wp.Name), "application/json", bytes.NewReader(updatePhpIniJson))
	if err != nil {
		fmt.Print(err)
	}
	defer resp.Body.Close()

	var phpIni models.PhpIni
	json.NewDecoder(resp.Body).Decode(&phpIni)

	ctx.JSON(200, phpIni)
}
