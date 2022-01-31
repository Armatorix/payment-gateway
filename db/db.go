package db

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Armatorix/payment-gateway/model"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type DB struct {
	*bun.DB
}

func Connect(dsn string) *DB {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	return &DB{bun.NewDB(sqldb, pgdialect.New())}
}

func (db *DB) VerifySecret(ctx context.Context, username, secret string) (bool, error) {
	m := model.Merchant{}
	err := db.NewSelect().
		Model(&m).
		Where("name = ?", username).
		Where("api_secret = ?", secret).
		Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (db *DB) SaveCreditCard(c model.CreditCard) (*model.CreditCard, error) {
	c.State = model.ActiveState
	// TODO
	return &model.CreditCard{}, nil
}
