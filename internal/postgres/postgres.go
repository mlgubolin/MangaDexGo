package postgres

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func ConnectToDatabase() *pgx.Conn {
	urlExample := "postgres://postgres:postgres@localhost:5432/mangadexgodb"
	conn, err := pgx.Connect(context.Background(), urlExample)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return conn
}

func CreateTable(conn *pgx.Conn, name string, table_info string) {
	RunQuery(conn, "CREATE TABLE IF NOT EXISTS "+name+" ( "+table_info+" );")
}

func RunQuery(conn *pgx.Conn, query string) pgx.Rows {
	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	return rows
}
