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

func UpdateModSecFirewall(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	siteId := ctx.Param("id")

	updateModSecFirewall := new(models.UpdateModSecFirewall)
	if err := ctx.ShouldBind(&updateModSecFirewall); err != nil {
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

	updateModSecFirewall.App = wp.Name

	updateModSecFirewallJSON, _ := json.Marshal(updateModSecFirewall)
	resp, err := http.Post(fmt.Sprintf("http://%s:8081/updateModsecurity", wp.IP), "application/json", bytes.NewReader(updateModSecFirewallJSON))
	if err != nil {
		fmt.Print(err.Error())
		ctx.AbortWithStatus(400)
		return
	}
	defer resp.Body.Close()

	err = site.UpdateModSecQuery(tx, wp.ID, updateModSecFirewall)
	if err != nil {
		fmt.Print(err.Error())
		ctx.AbortWithStatus(400)
		return
	}

	FirewallBySiteId(ctx)
}
