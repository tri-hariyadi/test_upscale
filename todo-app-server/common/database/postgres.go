package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
	"todo-app/common/config"
	"todo-app/common/logger"
)

type DBQuerier interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

func ConnectDB() *pgxpool.Pool {
	conf := config.AppConfig.Database
	ctx := context.Background()
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		conf.User, conf.Password, conf.Host, conf.Port, conf.Name)

	dbConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		logger.Logger.Fatalf("Error parsing database config: %v", err)
	}

	dbConfig.MaxConns = conf.MaxConns // sama seperti SetMaxOpenConns
	dbConfig.MinConns = conf.MinConns // mirip SetMaxIdleConns
	dbConfig.MaxConnIdleTime = time.Duration(conf.MaxConnIdleTime) * time.Minute
	dbConfig.MaxConnLifetime = time.Duration(conf.MaxConnLifetime) * time.Minute

	pool, err := pgxpool.NewWithConfig(ctx, dbConfig)
	if err != nil {
		logger.Logger.Fatalf("Error connecting to database: %v", err)
	}

	return pool
}
