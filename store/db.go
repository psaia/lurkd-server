package store

import (
	"database/sql"
	"fmt"
	"log"

	// Postgres...
	_ "github.com/lib/pq"
)

// DB contains the db singleton
var DB *sql.DB

// Connect with postgres
func Connect(pgUser string, pgPass string, pgHost string, pgDbName string) {
	driver, err := sql.Open(
		"postgres",
		fmt.Sprintf(
			"postgres://%s:%s@%s/%s?sslmode=disable",
			pgUser,
			pgPass,
			pgHost,
			pgDbName,
		),
	)

	if err != nil {
		log.Printf("DB Conn Fail: Host: %s, DB: %s, User: %s", pgHost, pgDbName, pgUser)
		panic(err)
	}

	if err = driver.Ping(); err != nil {
		log.Printf("DB driver ping fail.")
		panic(err)
	}

	DB = driver
}
