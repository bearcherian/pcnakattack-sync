package db

import (
	"bytes"
	"database/sql"
	"github.com/bearcherian/pcnakattackSync/config"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var connection *sql.DB

func GetClient() *sql.DB {
	cfg := config.GetConfig()
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

func Close() {
	if connection != nil {
		connection.Close()
	}
}
