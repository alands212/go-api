package database

import (
	"database/sql"

	"github.com/alands212/go-api/internal/logs"
	_ "github.com/go-sql-driver/mysql"
)

type MySqlClient struct {
	*sql.DB
}

func NewMySQLClient() *MySqlClient {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/loginserve")

	if err != nil {
		logs.Error("cannot create mysql client")
		panic(err)
	}

	err = db.Ping()

	if err != nil {

	}

	return &MySqlClient{db}
}
