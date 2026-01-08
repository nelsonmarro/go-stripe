// Package driver implements the database driver connection
package driver

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// ConnectSQL creates a new database connection
func ConnectSQL(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil
}
