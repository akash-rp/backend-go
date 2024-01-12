package site

import (
	"backend/db"
	"backend/models"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
)

func CheckIfSiteWithNameExists(server *models.ServerDetails, siteName string) (bool, error) {
	dbServerRow, err := db.DbConn.Query(context.Background(), "SELECT name from sites WHERE serverid = $1 AND name = $2", server.Id, siteName)
	if err != nil {
		return false, err
	}

	_, err = pgx.CollectOneRow(dbServerRow, pgx.RowToAddrOfStructByNameLax[models.SiteDetails])
	if err != nil {
		if err == pgx.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func GetSiteDetails(siteId string, userId string, tx pgx.Tx) (models.SiteDetails, error) {
	siteQuery, err := tx.Query(context.Background(), "SELECT sites.id, serverid, sites.name, php, staging, type, authentication, userid, \"user\", ip from sites JOIN servers ON sites.id = $1 AND sites.userid = $2 AND sites.serverid = servers.id", siteId, userId)
	if err != nil {
		fmt.Printf("%+v", err)
		return models.SiteDetails{}, err
	}

	site, err := pgx.CollectOneRow(siteQuery, pgx.RowToStructByNameLax[models.SiteDetails])
	if err != nil {
		//l.Printf("%+v", err)
		log.Print(err)
		return models.SiteDetails{}, err
	}

	return site, nil
}
