package sqlkit

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

type PGConfig struct {
	URL      string
	Debug    bool
	PoolSize int
}

type PGClient struct {
	*bun.DB
	Closed bool
}

func NewPGClient(ctx context.Context, conf *PGConfig) *PGClient {

	config, err := pgx.ParseConfig(conf.URL)
	if err != nil {
		panic("Failed to parse PostgreSQL URL")
	}
	config.PreferSimpleProtocol = true
	// config.Config.RuntimeParams["binary_parameters"] = "yes"

	sqldb := stdlib.OpenDB(*config)
	sqldb.SetMaxOpenConns(conf.PoolSize)

	db := bun.NewDB(sqldb, pgdialect.New())

	return &PGClient{
		DB:     db,
		Closed: false,
	}
}
