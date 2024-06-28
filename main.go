package main

import (
	"database/sql"
	"log"

	"github.com/support-sphere/support-sphere/internal/app"
)

func main() {
	connStr := "user=root dbname=support_sphere sslmode=disable password=root"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Could not ping the database: %v", err)
	}

	app.StartServer(db)
}
