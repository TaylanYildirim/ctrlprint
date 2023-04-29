package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func NewDB(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	go runHealthCheck(db, time.Minute*5)

	return db, nil
}

func runHealthCheck(db *sql.DB, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		err := db.Ping()
		if err != nil {
			log.Printf("Database health check failed: %v", err)
		} else {
			log.Println("Database health check passed")
		}
	}
}
