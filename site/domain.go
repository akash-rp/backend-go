package site

import (
	"backend/db"
	"backend/models"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	tld "github.com/jpillora/go-tld"
	"net/http"
	"strings"
)

func ModifyDomain(originalURL string) (models.ModifiedDomain, error) {
	var url string
	var isSubDomain bool
	routing := "none"

	httpUrl := EnsureHTTPProtocol(originalURL)

	parsedDomain, err := tld.Parse(httpUrl)
	if err != nil {
		return models.ModifiedDomain{}, err
	}

	if parsedDomain.Subdomain != "" {
		subDomains := strings.Split(parsedDomain.Subdomain, ".")
		if len(subDomains) == 1 {
			if subDomains[0] == "www" {
				url = fmt.Sprintf("%s.%s", parsedDomain.Domain, parsedDomain.TLD)
				routing = "www"
			} else {
				isSubDomain = true
				url = fmt.Sprintf("%s.%s.%s", subDomains[0], parsedDomain.Domain, parsedDomain.TLD)
			}
		} else {
			isSubDomain = true
			subDomains := strings.Join(subDomains, ".")
			url = fmt.Sprintf("%s.%s.%s", subDomains, parsedDomain.Domain, parsedDomain.TLD)
		}
	} else {
		url = fmt.Sprintf("%s.%s", parsedDomain.Domain, parsedDomain.TLD)
	}

	result := models.ModifiedDomain{
		Url:         url,
		Routing:     routing,
		IsSubDomain: isSubDomain,
	}

	return result, nil
}

func CheckIfDomainExists(serverId uuid.UUID, domain string) (bool, error) {

	dbServerRow, err := db.DbConn.Query(context.Background(), "SELECT url from domains WHERE serverid = $1 AND url = $2", serverId, domain)
	if err != nil {
		return false, err
	}

	_, err = pgx.CollectOneRow(dbServerRow, pgx.RowToAddrOfStructByNameLax[models.ModifiedDomain])
	if err != nil {
		if err == pgx.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func InsertDomainQuery(tx pgx.Tx, parsedDomain *models.ModifiedDomain, domainType int, ssl bool, wildcard bool, serverId uuid.UUID, siteId string) (*models.Domain, error) {
	resultQuery, err := tx.Query(context.Background(), "INSERT INTO domains(url, type, ssl, wildcard, subdomain, routing, siteid, serverid) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) returning id, url, type, ssl, wildcard, subdomain, routing, siteid;", parsedDomain.Url, domainType, ssl, wildcard, parsedDomain.IsSubDomain, parsedDomain.Routing, siteId, serverId)
	if err != nil {
		fmt.Print(err.Error())
		return &models.Domain{}, err
	}

	result, err := pgx.CollectOneRow(resultQuery, pgx.RowToAddrOfStructByNameLax[models.Domain])
	if err != nil {
		fmt.Print(err.Error())
		return &models.Domain{}, err
	}
	return result, nil
}

func DeleteDomainQuery(tx pgx.Tx, domainId string, siteId string) error {
	fmt.Print(domainId, siteId)
	_, err := tx.Exec(context.Background(), "DELETE FROM domains WHERE domains.id = $1 AND domains.siteid = $2 AND domains.type != $3", domainId, siteId, 1)
	if err != nil {
		fmt.Print(err.Error())
		return err
	}

	return nil
}

func GetDomainById(tx pgx.Tx, id string, siteId string) (models.Domain, error) {
	domainQuery, err := tx.Query(context.Background(), "SELECT id, url, subdomain, routing, wildcard, type, ssl FROM domains WHERE id = $1 AND siteid = $2", id, siteId)
	if err != nil {
		return models.Domain{}, err
	}

	domain, err := pgx.CollectOneRow(domainQuery, pgx.RowToStructByNameLax[models.Domain])
	if err != nil {
		fmt.Printf("%+v", err)
		return models.Domain{}, err
	}

	return domain, nil
}

func EnsureHTTPProtocol(url string) string {
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		return "http://" + url
	}
	return url
}

func UpdateDomainQuery(tx pgx.Tx, domain models.Domain, domainId uuid.UUID, siteId uuid.UUID) (models.Domain, error) {
	fmt.Print(domain)
	var result models.Domain
	row := tx.QueryRow(context.Background(), "UPDATE domains SET type = $1, ssl = $2, wildcard = $3, routing = $4 WHERE id = $5 AND siteid = $6 returning *;", domain.Type, domain.Ssl, domain.Wildcard, domain.Routing, domainId, siteId)
	err := row.Scan(&result)
	if err != nil {
		fmt.Print(err.Error())
		return models.Domain{}, err
	}

	return result, nil
}

func GetDomainsBySite(ctx *gin.Context, site models.SiteDetails) {
	tx, err := db.DbConn.Begin(context.Background())
	if err != nil {
		fmt.Print(err.Error())
		ctx.AbortWithStatus(400)
		return
	}
	defer tx.Commit(context.Background())

	domainsQuery, _ := tx.Query(context.Background(), "SELECT id, url, type, ssl, wildcard, subdomain, routing, siteid from domains WHERE siteid = $1", site.ID)
	if err != nil {
		fmt.Printf("%+v", err)
		ctx.AbortWithStatus(400)
		return
	}

	domains, err := pgx.CollectRows(domainsQuery, pgx.RowToStructByNameLax[models.Domain])
	if err != nil {
		fmt.Printf("%+v", err)
		ctx.AbortWithStatus(400)
		return
	}

	ctx.JSON(http.StatusOK, domains)
}

func GetPrimaryDomain(tx pgx.Tx, siteId string) (models.Domain, error) {
	domainQuery, err := tx.Query(context.Background(), "SELECT id, url, subdomain, routing, wildcard, type, ssl FROM domains WHERE type = 1 AND siteid = $1", siteId)
	if err != nil {
		return models.Domain{}, err
	}

	domain, err := pgx.CollectOneRow(domainQuery, pgx.RowToStructByNameLax[models.Domain])
	if err != nil {
		fmt.Printf("%+v", err)
		return models.Domain{}, err
	}

	return domain, nil
}
