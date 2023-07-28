package db

import (
	"database/sql"
	"fmt"
	"go-games-api/config"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	conf := config.New()
	db, err := sql.Open("mysql", ConnectionString(conf.Database))

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func ConnectionString(d config.DbConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", d.DbUser, d.DbPass, d.DbHost, d.DbPort, d.DbName)
}
