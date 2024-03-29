package controllers

import (
	"backend/db"
	"backend/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"net/http"
)

func GetBannedIps(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")

	dbServerRow, err := db.DbConn.Query(context.Background(), "SELECT ip from servers WHERE \"userId\" = $1 AND id = $2", userId, ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatus(404)
		return
	}
	serverDetails, err := pgx.CollectOneRow(dbServerRow, pgx.RowToAddrOfStructByNameLax[models.ServerDetails])

	if err != nil {
		ctx.AbortWithStatus(404)
		return
	}

	resp, err := http.Get(fmt.Sprintf("http://%s:8081/fail2ban/ip", serverDetails.IP))
	if err != nil {
		fmt.Print(err)
	}
	defer resp.Body.Close()

	var sshUsers models.SshUsers
	json.NewDecoder(resp.Body).Decode(&sshUsers)
	// result.Stats, _ = io.ReadAll(resp.Body)

	ctx.JSON(200, sshUsers)
}
