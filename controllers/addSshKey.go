package controllers

import (
	"backend/db"
	"backend/models"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"net/http"
	"time"
)

func AddSshKey(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	sshKey := new(models.SshKey)

	if err := ctx.ShouldBind(&sshKey); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	sshKey.Timestamp = time.Now().UnixMilli()

	dbServerRow, err := db.DbConn.Query(ctx, "SELECT ip from servers WHERE \"userId\" = $1 AND id = $2", userId, ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatus(404)
		return
	}

	serverDetails, err := pgx.CollectOneRow(dbServerRow, pgx.RowToAddrOfStructByNameLax[models.ServerDetails])
	if err != nil {
		ctx.AbortWithStatus(404)
		return
	}

	sshKeyJSON, _ := json.Marshal(sshKey)
	resp, err := http.Post(fmt.Sprintf("http://%s:8081/sshKey/add", serverDetails.IP), "application/json", bytes.NewReader(sshKeyJSON))
	if err != nil {
		fmt.Print(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errorMessage := new(models.Error)
		json.NewDecoder(resp.Body).Decode(&errorMessage)
		ctx.JSON(resp.StatusCode, errorMessage)
		return
	}

	var sshKeys []models.SshKey
	json.NewDecoder(resp.Body).Decode(&sshKeys)

	ctx.JSON(200, sshKeys)
}
