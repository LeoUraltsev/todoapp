package postgesql

import (
	"context"
	"fmt"
	"github.com/LeoUraltsev/todoapp/internal/config"
	"github.com/jackc/pgx/v5"
)

func NewClient(ctx context.Context, sc config.StorageConfig) (*pgx.Conn, error) {
	// urlExample := "postgres://username:password@localhost:5432/database_name"

	connectionURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", sc.Username, sc.Password, sc.Host, sc.Port, sc.Db)

	conn, err := pgx.Connect(ctx, connectionURL)
	if err != nil {
		return nil, fmt.Errorf("error connection to DB: %v", err)
	}

	return conn, err
}
