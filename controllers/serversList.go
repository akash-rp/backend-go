package controllers

import (
	"backend/db"
	"backend/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func ServersList(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	fmt.Print(userId)
	rows, err := db.DbConn.Query(ctx, "SELECT id, name, ip from servers WHERE \"userId\" = $1", userId)
	if err != nil {
		fmt.Print(err)
	}

	result, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.ServerList])
	if err != nil {
		fmt.Print(err)
		return
	}

	ctx.JSON(200, gin.H{
		"servers": result,
	})
}
