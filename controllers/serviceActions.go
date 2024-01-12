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
)

func ServiceAction(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	var serviceActionBody models.ServiceAction

	if err := ctx.ShouldBind(&serviceActionBody); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

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

	fmt.Print(serviceActionBody)
	fmt.Print(serviceActionBody)
	serviceActionBodyJSON, _ := json.Marshal(serviceActionBody)
	resp, err := http.Post(fmt.Sprintf("http://%s:8081/service", serverDetails.IP), "application/json", bytes.NewBuffer(serviceActionBodyJSON))
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

	var serviceStatus models.ServiceStatus
	fmt.Print(resp.StatusCode)
	json.NewDecoder(resp.Body).Decode(&serviceStatus)

	ctx.JSON(200, gin.H{
		"services": serviceStatus,
	})

}
