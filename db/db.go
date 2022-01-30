package db

import (
	"context"
	"database/sql"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type TxContext struct {
	echo.Context
	Tx    bun.Tx
	TxCtx context.Context
}

type DB struct {
	*bun.DB
}

func Connect(dsn string) *DB {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	return &DB{bun.NewDB(sqldb, pgdialect.New())}
}

func (db *DB) VerifySecret(username, secret string) (bool, error) {
	// TODO db.NewSelect().Model()
	return true, nil
}
