package db

import (
	"api/src/config"
	"database/sql"

	_ "github.com/lib/pq" // add driver
)

func Connect() (*sql.DB, error) {

	dataBase, err := sql.Open("postgres", config.JdbcUrl)
	if err != nil {
		return nil, err
	}

	if err = dataBase.Ping(); err != nil {
		dataBase.Close()
		return nil, err
	}

	return dataBase, nil
}
