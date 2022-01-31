package db

import (
	"database/sql"

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

func (db *DB) VerifySecret(username, secret string) (bool, error) {
	// TODO db.NewSelect().Model()
	return true, nil
}

func (db *DB) SaveCreditCard(c model.CreditCard) (*model.CreditCard, error) {
	// TODO
	return &model.CreditCard{}, nil
}
