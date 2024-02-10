package database

import (
	"context"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5"
)

// DataBase представляет соединение с базой данных.
type DataBase struct {
	*pgx.Conn
}

// New создает новое соединение с базой данных и выполняет миграции.
func New() (*DataBase, error) {
	connURL := createConnectionURL()
	connConfig, _ := pgx.ParseConfig(connURL)

	conn, err := pgx.ConnectConfig(context.Background(), connConfig)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to the database: %v\n", err)
	}

	log.Printf("Postgres connected")

	db := DataBase{
		Conn: conn,
	}

	m, err := migrate.New(
		"file://internal/database/migrations",
		connURL,
	)
	if err != nil {
		return nil, fmt.Errorf("unable to migrate the database: %v\n", err)
	}

	err = m.Up()
	if err != nil {
		log.Printf("Unable to migrate the database: %v\n", err)
	}

	return &db, nil
}

// Close закрывает соединение с базой данных.
func (db *DataBase) Close() {
	if db.Conn != nil {
		db.Conn.Close(context.Background())
	}
}

// createConnectionURL создает URL для подключения к базе данных.
func createConnectionURL() string {
	host := "127.0.0.1"
	port := 5432
	user := "postgres"
	password := "123456"
	database := "app"

	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port, database)
}
