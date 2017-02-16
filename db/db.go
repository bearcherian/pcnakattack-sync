package db

import (
	"github.com/bearcherian/pcnakattackSync/config"
	"bytes"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var connection *sql.DB

func GetClient(cfg config.Config) *sql.DB {
	if connection == nil {
		var dsn bytes.Buffer
		dsn.WriteString(cfg.Database.Username)
		dsn.WriteString(":")
		dsn.WriteString(cfg.Database.Password)
		dsn.WriteString("@tcp(")
		dsn.WriteString(cfg.Database.Hostname)
		dsn.WriteString(")/")
		dsn.WriteString(cfg.Database.Database)
		dsn.WriteString("?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=true")
		var err error
		connection, err = sql.Open("mysql", dsn.String())
		if err != nil {
			log.Fatal(err)
		}
	}

	return connection
}
