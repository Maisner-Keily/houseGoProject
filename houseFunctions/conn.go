package houseFunctions

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
)

func getConn(ctx context.Context) pgx.Conn {
	url := fmt.Sprintf("postgres://%s:%s@postgres:%s/%s", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"))
	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return *conn
}
