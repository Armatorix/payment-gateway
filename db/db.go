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
	_, err := db.NewInsert().
		Model(&c).Exec(ctx, &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (db *DB) CaptureSpending(ctx context.Context, mID uuid.UUID, c model.CardAction) (*model.CreditCard, error) {
	var card model.CreditCard
	err := db.RunInTx(ctx,
		&sql.TxOptions{Isolation: sql.LevelSerializable},
		func(ctx context.Context, tx bun.Tx) error {
			// get card info (with specified merchantID to verify access permission;
			//               ;with currency to verify if it is proper one)
			// get all  actions
			// verify if card is still active (not refunded, not voided) and with funds(new action has lower amount than left)
			// insert action
			return nil
		})
	if err != nil {
		return nil, err
	}
	return &card, nil
}
