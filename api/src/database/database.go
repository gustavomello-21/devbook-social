package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gustavomello-21/devbook/api/src/config"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.ConnectionUrl)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
