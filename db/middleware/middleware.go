package middleware

import (
	"context"
	"database/sql"

	"github.com/Armatorix/payment-gateway/db"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/uptrace/bun"
)

func SerializableTxMiddleware(sqlDB *db.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return sqlDB.RunInTx(
				c.Request().Context(),
				&sql.TxOptions{Isolation: sql.LevelSerializable},
				func(ctx context.Context, tx bun.Tx) error {
					return next(db.TxContext{Context: c, Tx: tx, TxCtx: ctx})
				})
		}
	}
}

func AuthMiddleware(sqlDB *db.DB) echo.MiddlewareFunc {
	return middleware.BasicAuth(func(username, secret string, c echo.Context) (bool, error) {
		return sqlDB.VerifySecret(username, secret)
	})
}
