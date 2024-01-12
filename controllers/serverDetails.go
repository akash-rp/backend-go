package controllers

import (
	"backend/db"
	"backend/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"net/http"
)

func ServerDetails(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	rows, err := db.DbConn.Query(ctx, "SELECT id, name, ip from servers WHERE \"userId\" = $1 AND id = $2", userId, ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatus(404)
		return
	}

	result, err := pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByNameLax[models.ServerDetails])
	if err != nil {
		ctx.AbortWithStatus(404)
		return
	}

	resp, err := http.Get(fmt.Sprintf("http://%s:8081/serverstats", result.IP))
	if err != nil {
		fmt.Print(err)
	}

	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&result.Stats)

	ctx.JSON(200, result)
}
