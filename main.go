package main

import (
	"Reminder/internal/db/postgres"
	"context"
)

func main() {
	ctx := context.Background()
	pg := postgres.New(ctx)

}
