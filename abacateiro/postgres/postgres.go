package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	Pool *pgxpool.Pool
}

func NewDB(connString string) (*DB, error) {
	pool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	err = pool.Ping(context.Background())
	if err != nil {
		return nil, fmt.Errorf("unable to ping database: %v", err)
	}

	return &DB{Pool: pool}, nil
}

func (db *DB) Close() {
	db.Pool.Close()
}

func (db *DB) Query(query string, args ...interface{}) error {
	rows, err := db.Pool.Query(context.Background(), query, args...)
	if err != nil {
		return fmt.Errorf("query error: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var result string // Adjust the type based on your query
		err := rows.Scan(&result)
		if err != nil {
			return fmt.Errorf("scan error: %v", err)
		}
		fmt.Println(result)
	}

	return rows.Err()
}
