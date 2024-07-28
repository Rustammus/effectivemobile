package crud

import (
	"EffectiveMobile/internal/config"
	"EffectiveMobile/pkg/client/postgres"
	"EffectiveMobile/pkg/logging"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

var connPool *pgxpool.Pool

func init() {
	logger := logging.GetLogger()

	conf := config.GetConfig()
	pool, err := postgres.NewPool(context.TODO(), conf.Storage)
	if err != nil {
		logger.Fatalf("Can't crate connection Pool. Abort start app. \n Error: %s", err.Error())
	}
	connPool = pool
}

func GetPool() *pgxpool.Pool {
	return connPool
}
