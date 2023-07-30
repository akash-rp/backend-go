package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(ctx *gin.Context) {
	client := http.Client{}
	req, err := http.NewRequest("GET", KARTOS_URL+"sessions/whoami", nil)

	if err != nil {
		ctx.AbortWithStatusJSON(403, gin.H{})
		return
	}

	req.Header = ctx.Request.Header

	res, err := client.Do(req)

	if err != nil || res.StatusCode != 200 {
		log.Print(res)
		ctx.AbortWithStatusJSON(403, gin.H{})
		return
	}
	ctx.Set("userId", res.Header.Get("X-Kratos-Authenticated-Identity-Id"))
	ctx.Next()
}
