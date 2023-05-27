package config

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

var Conn *sql.DB

func OpenConn() {
	cfg := mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASS"),
		Addr:                 os.Getenv("DB_HOST"),
		DBName:               os.Getenv("DB_NAME"),
		Net:                  "tcp",
		AllowNativePasswords: true,
	}

	var err error

	Conn, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	err = Conn.Ping()
	if err != nil {
		log.Fatal(err)
	}
}
