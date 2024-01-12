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

func PhpIni(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	siteId := ctx.Param("id")

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

	resp, err := http.Get(fmt.Sprintf("http://%s:8081/getPHPini/%s", wp.IP, wp.Name))
	if err != nil {
		fmt.Print(err)
	}
	defer resp.Body.Close()

	var phpIni models.PhpIni
	json.NewDecoder(resp.Body).Decode(&phpIni)

	ctx.JSON(200, phpIni)
}
