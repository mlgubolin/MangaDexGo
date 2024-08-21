package mangadex

import (
	"mangadexgo/internal/postgres"

	"github.com/jackc/pgx/v5"
)

func MangaDex() *pgx.Conn {
	return postgres.ConnectToDatabase()

}
