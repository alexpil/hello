package db

import (
	"context"
	"time"

	pgx "github.com/jackc/pgx/v4"
)

func Connect(source string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), source)
	if err != nil {
		return nil, err
	}

	// Try ping to check connection
	attempts := 10
	for i := 0; i < attempts; i++ {
		err = conn.Ping(context.Background())
		if err == nil {
			break
		}
		time.Sleep(5 * time.Second)
	}
	if err != nil {
		return nil, err
	}

	return conn, nil
}
