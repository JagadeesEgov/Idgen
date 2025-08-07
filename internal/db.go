package internal

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"idgen/internal/migrations"

	_ "github.com/lib/pq"
)

func InitDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open db: %w", err)
	}
	// Retry ping for up to 10 seconds
for i := 0; i < 10; i++ {
    if err := db.Ping(); err == nil {
        break
    }
    log.Println("Waiting for DB to be ready...")
    time.Sleep(1 * time.Second)
}


	// Run database migrations
	if err := migrations.RunMigrations(db); err != nil {
		return nil, fmt.Errorf("failed to run migrations: %w", err)
	}

	return db, nil
}
