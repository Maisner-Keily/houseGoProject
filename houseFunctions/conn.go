package houseFunctions

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
)

func getConn(ctx context.Context) pgx.Conn {
	url := "postgres://postgres:admin@postgres:5432/mydockerhouse"
	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return *conn
}
