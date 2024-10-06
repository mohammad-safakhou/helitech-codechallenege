package utils

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"time"
)

func PostgresConnection(host, port, user, pass, database, sslmode string, maxOpenConns, maxIdleConns int, timeout time.Duration) (*sql.DB, error) {
	connString := PostgresURI(host, port, user, pass, database, sslmode)
	log.Println("postgres options -> " + connString)
	conn, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, fmt.Errorf("error in openning postgres connection: %w", err)
	}

	conn.SetMaxOpenConns(maxOpenConns)
	conn.SetMaxIdleConns(maxIdleConns)

	dbContext, _ := context.WithTimeout(context.Background(), timeout)
	err = conn.PingContext(dbContext)
	if err != nil {
		return nil, fmt.Errorf("error in pinging postgres database: %w", err)
	}
	return conn, nil
}

func PostgresURI(host, port, user, pass, database, sslmode string) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", user, pass, host, port, database, sslmode)
}
