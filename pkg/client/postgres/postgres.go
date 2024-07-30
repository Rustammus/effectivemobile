package postgres

import (
	"EffectiveMobile/internal/config"
	"EffectiveMobile/pkg/logging"
	"context"
	"errors"
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
			logger.Error(err.Error())
			maxAttempts--
			time.Sleep(time.Second)
			continue
		}
		return connect, err
	}
	return nil, errors.New("limit connection try exceeded")
}

func NewPool(ctx context.Context, cf config.Storage) (pool *pgxpool.Pool, err error) {
	logger := logging.GetLogger()
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", cf.Username, cf.Password, cf.Host, cf.Port, cf.Database)
	maxAttempts := 5

	ctxPool, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	pool, err = pgxpool.New(ctxPool, dsn)
	if err != nil {
		return nil, err
	}

	ctxPing, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	for ; maxAttempts > 0; maxAttempts-- {
		if err = pool.Ping(ctxPing); err != nil {
			logger.Error(err.Error())
			time.Sleep(time.Second)
			continue
		}
		logger.Infof("Connected to database %s", cf.Database)
		return pool, nil
	}

	return nil, errors.New("limit connection try exceeded")
}
