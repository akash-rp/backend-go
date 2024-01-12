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

// Need refactor to use one function and get method
func ServerHealthByName(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")

	dbServerRow, err := db.DbConn.Query(ctx, "SELECT ip from servers WHERE \"userId\" = $1 AND id = $2", userId, ctx.Param("id"))

	if err != nil {
		ctx.AbortWithStatus(404)
		return
	}
	result, err := pgx.CollectOneRow(dbServerRow, pgx.RowToAddrOfStructByNameLax[models.ServerDetails])

	if err != nil {
		ctx.AbortWithStatus(404)
		return
	}

	resp, err := http.Get(fmt.Sprintf("http://%s:8081/metrics", result.IP))
	if err != nil {
		fmt.Print(err)
	}
	defer resp.Body.Close()

	var health models.ServerHealth
	json.NewDecoder(resp.Body).Decode(&health)
	// result.Stats, _ = io.ReadAll(resp.Body)

	ctx.JSON(200, gin.H{
		"health": health,
	})
}
