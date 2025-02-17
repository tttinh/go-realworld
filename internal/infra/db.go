package infra

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func ConnectDB() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), "postgres://tinhtt:tinhtt@localhost:5432/conduit?sslmode=disable")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	return conn
}

func CloseDB(conn *pgx.Conn) {
	err := conn.Close(context.Background())
	if err != nil {
		log.Fatalf("Unable to close connection: %v\n", err)
	}
}
