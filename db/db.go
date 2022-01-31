package db

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Armatorix/payment-gateway/model"
	"github.com/google/uuid"
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

func (db *DB) VerifySecret(ctx context.Context, username, secret string) (bool, uuid.UUID, error) {
	m := model.Merchant{}
	err := db.NewSelect().
		Model(&m).
		Where("name = ?", username).
		Where("api_secret = ?", secret).
		Scan(ctx, &m)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, uuid.Nil, nil
		}
		return false, uuid.Nil, err
	}
	return true, m.ID, nil
}

func (db *DB) SaveCreditCard(ctx context.Context, c model.CreditCard) (*model.CreditCard, error) {
	c.CardState = model.ActiveState
	_, err := db.NewInsert().
		Model(&c).Exec(ctx, &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
