package internal

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func InitDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open db: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping db: %w", err)
	}

	// Ensure idgen_templates table exists
	tableSQL := `
	CREATE TABLE IF NOT EXISTS idgen_templates (
		id          VARCHAR(64) PRIMARY KEY,
		config      JSONB NOT NULL,
		created_at  BIGINT,
		created_by  VARCHAR(64)
	);
	`
	_, err = db.Exec(tableSQL)
	if err != nil {
		return nil, fmt.Errorf("failed to create idgen_templates table: %w", err)
	}

	return db, nil
} 