package config

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
)

func ExecMigrations() {
	connection := Conn

	driver, _ := mysql.WithInstance(connection, &mysql.Config{})

	var err error

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"mysql",
		driver,
	)

	if err != nil {
		log.Fatal(err)
	}

	m.Up()

}
