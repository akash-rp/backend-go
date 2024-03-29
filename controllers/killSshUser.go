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

func KillSshUser(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	sshUser := new(models.KillSshUser)

	if err := ctx.ShouldBind(&sshUser); err != nil {
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

	sshUserJSON, _ := json.Marshal(sshUser)
	resp, err := http.Post(fmt.Sprintf("http://%s:8081/ssh/kill", serverDetails.IP), "application/json", bytes.NewBuffer(sshUserJSON))
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

	var sshUsers models.SshUsers

	json.NewDecoder(resp.Body).Decode(&sshUsers)
	// result.Stats, _ = io.ReadAll(resp.Body)

	ctx.JSON(200, sshUsers)
}
