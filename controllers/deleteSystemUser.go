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

func DeleteSystemUser(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	user := new(models.DeleteSystemUser)

	if err := ctx.ShouldBind(&user); err != nil {
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

	userJSON, _ := json.Marshal(user)
	resp, err := http.Post(fmt.Sprintf("http://%s:8081/users/delete", serverDetails.IP), "application/json", bytes.NewReader(userJSON))
	if err != nil {
		ctx.AbortWithStatus(400)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errorMessage := new(models.Error)
		json.NewDecoder(resp.Body).Decode(&errorMessage)
		ctx.JSON(resp.StatusCode, errorMessage)
		return
	}

	systemUsers := new(models.SystemUsers)
	json.NewDecoder(resp.Body).Decode(&systemUsers)
	ctx.JSON(http.StatusOK, systemUsers)
}
