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

func ManageUfwRule(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	var ufwRule interface{}
	var ufwRuleJSON []byte
	fmt.Print(ctx.Request.Method)

	switch ctx.Request.Method {
	case "POST":
		ufwRule = new(models.AddUfwRule)
	case "DELETE":
		ufwRule = new(models.DeleteUfwRule)
	default:
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := ctx.ShouldBind(ufwRule); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	dbServerRow, err := db.DbConn.Query(ctx, "SELECT ip from servers WHERE \"userId\" = $1 AND id = $2", userId, ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}
	serverDetails, err := pgx.CollectOneRow(dbServerRow, pgx.RowToAddrOfStructByNameLax[models.ServerDetails])

	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	client := http.Client{}
	req, err := http.NewRequest(ctx.Request.Method, fmt.Sprintf("http://%s:8081/ufw", serverDetails.IP), bytes.NewBuffer(ufwRuleJSON))
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errorMessage := new(models.Error)
		json.NewDecoder(resp.Body).Decode(&errorMessage)
		ctx.JSON(resp.StatusCode, errorMessage)
		return
	}

	ufwRulesList := new(models.UfwRules)
	json.NewDecoder(resp.Body).Decode(&ufwRulesList)

	ctx.JSON(http.StatusOK, ufwRulesList)
}
