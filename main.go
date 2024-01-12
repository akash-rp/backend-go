package main

import (
	"backend/db"
	"context"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid/v5"
	pgxuuid "github.com/jackc/pgx-gofrs-uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	fmt.Printf("sfwe")
	var err error
	db.DbConn, err = pgxpool.New(context.Background(), "postgres://akash-rp:Z9fMpYVQH0tI@ep-sparkling-glade-138972.ap-southeast-1.aws.neon.tech/backend")
	db.DbConn.Config().AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		pgxuuid.Register(conn.TypeMap())
		return nil
	}
	//tx, _ := db.DbConn.Begin(context.Background())
	//defer db.DbConn.Close()
	if err != nil {
		fmt.Print(err)
	}
	route := gin.Default()
	uid, _ := uuid.NewV4()
	// route.Use(Options)
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://127.0.0.1:3000"}
	corsConfig.AllowCredentials = true
	route.Use(cors.New(corsConfig))
	route.GET("/test", func(ctx *gin.Context) {
		fmt.Print(ctx.Request.Header.Get("X-user-id"))
		ctx.JSON(200, gin.H{
			"uuid": uid,
		})
	})
	addRoutes(route)
	route.Run("127.0.0.1:4288")
}
