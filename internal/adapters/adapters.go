package adapters

import (
	"context"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5"
	"github.com/tinhtt/go-realworld/internal/config"
)

func ConnectDB(c config.Config) (*pgx.Conn, error) {
	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		c.Database.Username,
		c.Database.Password,
		c.Database.Host,
		c.Database.Port,
		c.Database.Name,
	)
	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		return nil, err
	}

	return conn, err
}

func CloseDB(conn *pgx.Conn) {
	conn.Close(context.Background())
}

func Migrate(c config.Config) error {
	path := fmt.Sprintf("file://%s", c.Migration.Path)
	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		c.Database.Username,
		c.Database.Password,
		c.Database.Host,
		c.Database.Port,
		c.Database.Name,
	)

	m, err := migrate.New(path, dbURL)
	if err != nil {
		return fmt.Errorf("fail to create migration: %v", err)
	}

	if c.Migration.Fresh {
		if err := m.Down(); err != nil {
			if err != migrate.ErrNoChange {
				return fmt.Errorf("fail to migrate down: %v", err)
			}
		}
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("fail to migrate up: %v", err)
	}

	return nil
}
