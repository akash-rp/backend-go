package site

import (
	"backend/models"
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func InsertFirewallQuery(tx pgx.Tx, siteId uuid.UUID) error {
	_, err := tx.Exec(context.Background(), "INSERT INTO firewall(siteid) VALUES ($1)", siteId)
	if err != nil {
		fmt.Print(err.Error())
		return err
	}

	return nil
}

func UpdateSevenGQuery(tx pgx.Tx, siteId uuid.UUID, sevenG *models.UpdateSevenGFirewall) error {
	_, err := tx.Exec(context.Background(), "UPDATE firewall SET \"sevenG\" = $1, \"sevenG_disabled\" = $2 WHERE siteid = $3", sevenG.Enabled, sevenG.Disable, siteId)
	if err != nil {
		fmt.Print(err.Error())
		return err
	}

	return nil
}

func UpdateModSecQuery(tx pgx.Tx, siteId uuid.UUID, modSec *models.UpdateModSecFirewall) error {
	_, err := tx.Exec(context.Background(), "UPDATE firewall SET modsecurity = $1, \"modsecurity_anomalyThreshold\" = $2, \"modsecurity_paranoiaLevel\" = $3 WHERE siteid = $4", modSec.Enabled, modSec.AnomalyThreshold, modSec.ParanoiaLevel, siteId)
	if err != nil {
		fmt.Print(err.Error())
		return err
	}

	return nil
}
