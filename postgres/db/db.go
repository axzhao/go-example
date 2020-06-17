package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type DB struct {
	db *sql.DB
}

func (db *DB) Query(data interface{}, sql string, args ...interface{}) error {
	rows, err := db.db.Query(sql, args...)
	if err != nil {
		return err
	}

	return nil
}
