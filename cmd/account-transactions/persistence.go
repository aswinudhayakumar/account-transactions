package main

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// NewDBConnection returns a new instance of Database object
func NewDBConnection(ctx context.Context, config env) (*sqlx.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		config.DBUser,
		config.DBPassword,
		config.DBName,
		config.DBHost,
		config.DBPort,
		config.SSLMode,
	)

	db, err := sqlx.ConnectContext(ctx, "postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
