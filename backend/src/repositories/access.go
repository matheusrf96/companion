package repositories

import (
	"database/sql"
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

func (repo Access) getTableName() (string, string, string) {
	return pq.QuoteIdentifier(fmt.Sprintf("accesses_%s", time.Now().Format("20060102"))),
		pq.QuoteLiteral(time.Now().Format("2006-01-02")),
		pq.QuoteLiteral(time.Now().Add(time.Hour * 24).Format("2006-01-02"))
}

func (repo Access) createDayTable() error {
	tableName, beginDate, endDate := repo.getTableName()

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
	return nil
}
