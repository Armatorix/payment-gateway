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

func DBSerializableTxMiddleware(db *bun.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return db.RunInTx(
				c.Request().Context(),
				&sql.TxOptions{Isolation: sql.LevelSerializable},
				func(ctx context.Context, tx bun.Tx) error {
					return next(TxContext{c, tx, ctx})
				})
		}
	}
}

func Connect(dsn string) *bun.DB {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	return bun.NewDB(sqldb, pgdialect.New())
}
