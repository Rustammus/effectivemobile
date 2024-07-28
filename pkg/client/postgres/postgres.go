package postgres

import (
	"EffectiveMobile/internal/config"
	"EffectiveMobile/pkg/logging"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

func NewClient(ctx context.Context, cf config.Storage) (connect *pgx.Conn, err error) {
	logger := logging.GetLogger()

	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", cf.Username, cf.Password, cf.Host, cf.Port, cf.Database)
	maxAttempts := 5

	for maxAttempts > 0 {
		ctxT, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		connect, err = pgx.Connect(ctxT, dsn)
		cancel()
		if err != nil {
			logger.Error("Failed to connect to database. Try again...")
			maxAttempts--
			time.Sleep(time.Second)
			continue
		}
		return connect, err
	}
	return nil, err
}

func NewPool(ctx context.Context, cf config.Storage) (pool *pgxpool.Pool, err error) {
	logger := logging.GetLogger()
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", cf.Username, cf.Password, cf.Host, cf.Port, cf.Database)
	maxAttempts := 5

	logger.Info("Connection to database")
	ctxPool, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	pool, err = pgxpool.New(ctxPool, dsn)
	if err != nil {
		logger.Error(err.Error())
		logger.Fatalln("Failed to create connection pool. Abort start app.")
	}

	ctxPing, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	for ; maxAttempts > 0; maxAttempts-- {
		if err = pool.Ping(ctxPing); err != nil {
			logger.Error(err.Error())
			time.Sleep(time.Second)
			continue
		}
		return pool, nil
	}
	logger.Fatalln("Failed to connect to database. Abort start app.")

	return nil, nil
}
