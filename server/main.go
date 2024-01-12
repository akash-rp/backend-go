package server

import (
	"backend/db"
	"backend/models"
	"context"
	"github.com/jackc/pgx/v5"
)

func GetServerDetails(serverId string, userId string) (*models.ServerDetails, error) {
	rows, err := db.DbConn.Query(context.Background(), "SELECT id, name, ip from servers WHERE \"userId\" = $1 AND id = $2", userId, serverId)
	if err != nil {
		return &models.ServerDetails{}, nil
	}

	result, err := pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByNameLax[models.ServerDetails])
	if err != nil {
		return &models.ServerDetails{}, nil
	}

	return result, nil
}
