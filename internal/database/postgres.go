package database

import (
	"context"
	"embed"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

//go:embed migrations/*.sql
var fs embed.FS

type ConnConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

func (c ConnConfig) ToString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", c.User, c.Password, c.Host, c.Port, c.Database)
}

type DataBase interface {
	Exec(ctx context.Context, sql string, arguments ...any) (commandTag pgconn.CommandTag, err error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

// New создает новое соединение с базой данных и выполняет миграции.
func New(connConfig ConnConfig) (*pgx.Conn, error) {
	config, err := pgx.ParseConfig(connConfig.ToString())
	if err != nil {
		return nil, fmt.Errorf("could not create ConnConfig: %v", err)
	}

	conn, err := pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to the database: %v", err)
	}

	log.Printf("Postgres connected successful")

	migrations, err := iofs.New(fs, "migrations")
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithSourceInstance(
		"iofs",
		migrations,
		connConfig.ToString(),
	)
	if err != nil {
		return nil, fmt.Errorf("unable to migrate the database: %v", err)
	}

	err = m.Up()
	if err != nil {
		log.Printf("Unable to migrate the database: %v\n", err)
	}

	return conn, nil
}

func WithConn() ConnConfig {
	return ConnConfig{
		Host:     "127.0.0.1",
		Port:     5432,
		User:     "postgres",
		Password: "123456",
		Database: "app",
	}
}

func WithTestConn() ConnConfig {
	return ConnConfig{
		Host:     "127.0.0.1",
		Port:     5433,
		User:     "postgres",
		Password: "123456",
		Database: "test",
	}
}
