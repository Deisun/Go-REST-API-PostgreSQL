package database

import (
	"database/sql"
	"fmt"
	"os"
)

var DB *sql.DB

func ConnectDB() {

	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	DB = db
}
