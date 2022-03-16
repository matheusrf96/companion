package repositories

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/lib/pq"
	"github.com/matheusrf96/go-webserver/backend/src/models"
)

type Access struct {
	db *sql.DB
}

func NewAccessRepository(db *sql.DB) *Access {
	return &Access{db}
}

func getTableName() (string, string, string) {
	return pq.QuoteIdentifier(fmt.Sprintf("accesses_%s", time.Now().Format("20060102"))),
		pq.QuoteLiteral(time.Now().Format("2006-01-02")),
		pq.QuoteLiteral(time.Now().Add(time.Hour * 24).Format("2006-01-02"))
}

func (repo Access) createDayTable() error {
	tableName, beginDate, endDate := getTableName()

	statement, err := repo.db.Prepare(fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s PARTITION OF accesses
			FOR VALUES FROM (%s) TO (%s)
	`, tableName, beginDate, endDate))
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec()
	if err != nil {
		return err
	}

	return nil
}

func (repo Access) Save(access models.Access) error {
	err := repo.createDayTable()
	if err != nil {
		return err
	}

	tableName, _, _ := getTableName()

	statement, err := repo.db.Prepare(fmt.Sprintf(`
		INSERT INTO %s (
			uuid
			, source_id
			, utm_source
			, utm_medium
			, tags
			, referrer
			, cookie
			, user_agent
			, query
			, device
			, os
			, browser
			, screen
			, navigator
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9,
			$10, $11, $12, $13, $14
		)
	`, tableName))
	if err != nil {
		return err
	}
	defer statement.Close()

	screenData, err := json.Marshal(access.Screen)
	if err != nil {
		return err
	}

	navigatorData, err := json.Marshal(access.Navigator)
	if err != nil {
		return err
	}

	_, err = statement.Exec(
		&access.Uuid, &access.DetailData.SourceId, &access.DetailData.UtmSource,
		&access.DetailData.UtmMedium, pq.Array(&access.DetailData.Tags),
		&access.Referrer, &access.Cookie, &access.UserAgent, &access.Query,
		&access.Device, &access.OS, &access.Browser, screenData, navigatorData,
	)
	if err != nil {
		return err
	}

	return nil
}
