package databases

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func GetPsqlClient(uri string) *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), uri)
	if err != nil {
		panic(err)
	}

	return conn
}
