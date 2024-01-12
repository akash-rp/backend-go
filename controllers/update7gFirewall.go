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

func UpdateSevenGFirewall(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	siteId := ctx.Param("id")

	updateSevenGFirewall := new(models.UpdateSevenGFirewall)
	if err := ctx.ShouldBind(&updateSevenGFirewall); err != nil {
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

	updateSevenGFirewall.App = wp.Name
	updateSevenGFirewall.User = wp.User

	updateSevenGFirewallJSON, _ := json.Marshal(updateSevenGFirewall)
	resp, err := http.Post(fmt.Sprintf("http://%s:8081/update7G", wp.IP), "application/json", bytes.NewReader(updateSevenGFirewallJSON))
	if err != nil {
		fmt.Print(err)
		ctx.AbortWithStatus(400)
		return
	}
	defer resp.Body.Close()

	err = site.UpdateSevenGQuery(tx, wp.ID, updateSevenGFirewall)
	if err != nil {
		fmt.Print(err.Error())
		ctx.AbortWithStatus(400)
		return
	}

	FirewallBySiteId(ctx)
}
