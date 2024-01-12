package middlewares

import (
	"backend/db"
	"backend/site"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

func SiteVerify(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	siteId := ctx.Param("id")
	tx, err := db.DbConn.Begin(ctx)
	if err != nil {
		fmt.Print(err.Error())
		ctx.AbortWithStatus(400)
		return
	}
	defer tx.Commit(ctx)
	//TODO
	//Need to verify if the user has access to the site by checking db with userId and siteId if not reject with 404
	wp, err := site.GetSiteDetails(siteId, userId.(string), tx)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"message": "Site not found.",
		})
	}
	wpJSON, err := json.Marshal(wp)
	ctx.Set("site", wpJSON)
	ctx.Next()
}
