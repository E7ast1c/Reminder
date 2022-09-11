package postgres

import (
	"Reminder/pkg/empty"
	"context"
	log "github.com/apex/log"
	"github.com/jackc/pgx/v4"
)

func New(ctx context.Context) *pgx.Conn {
	uri := empty.EnvMustNotEmpty("PG_URI")
	conn, err := pgx.Connect(ctx, uri)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	return conn
}
