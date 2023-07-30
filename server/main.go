package server

import (
	"backend/db"
	"context"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5"
	"net/http"

	"backend/models"

	"github.com/gin-gonic/gin"
)

func getServersList(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	rows, error := db.DbConn.Query(context.Background(), "SELECT id, name, ip from servers WHERE \"userId\" = $1", userId)
	if error != nil {
		fmt.Print(error)
	}

	result, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.ServerList])
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Print(ctx.Get("userId"))
	ctx.JSON(200, gin.H{
		"servers": result,
	})

}

// Need to implement db request to get sites list for this server
func getServerDetails(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	rows, error := db.DbConn.Query(context.Background(), "SELECT id, name, ip from servers WHERE \"userId\" = $1 AND id = $2", userId, ctx.Param("id"))
	if error != nil {
		ctx.AbortWithStatus(404)
		return
	}
	result, err := pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByNameLax[models.ServerDetails])
	if err != nil {
		ctx.AbortWithStatus(404)
		return
	}
	resp, err := http.Get(fmt.Sprintf("http://%s:8081/serverstats", result.IP))
	// fmt.Print(stats.)
	if err != nil {
		fmt.Print(err)
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&result.Stats)
	// result.Stats, _ = io.ReadAll(resp.Body)

	ctx.JSON(200, result)

}

func getServerHealth(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")

	dbServerRow, err := db.DbConn.Query(context.Background(), "SELECT ip from servers WHERE \"userId\" = $1 AND id = $2", userId, ctx.Param("id"))

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

// Need refactor in agnet to fetch by post request
func fetchServerHealthByName(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")

	dbServerRow, err := db.DbConn.Query(context.Background(), "SELECT ip from servers WHERE \"userId\" = $1 AND id = $2", userId, ctx.Param("id"))

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

func getServicesStatus(ctx *gin.Context) {
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

	resp, err := http.Get(fmt.Sprintf("http://%s:8081/service/status", serverDetails.IP))
	if err != nil {
		fmt.Print(err)
	}
	defer resp.Body.Close()

	var serviceStatus models.ServiceStatus
	json.NewDecoder(resp.Body).Decode(&serviceStatus)
	// result.Stats, _ = io.ReadAll(resp.Body)

	ctx.JSON(200, gin.H{
		"services": serviceStatus,
	})
}
