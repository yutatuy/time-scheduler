package testutil

import (
	"context"
	"go-app/src/domain/shared"

	"gorm.io/gorm"
)

type tx struct {
	db *gorm.DB
}

func NewTestTransaction(db *gorm.DB) shared.Transaction {
	return &tx{db: db}
}

var txKey = struct{}{}

func (t *tx) DoInTx(ctx context.Context, f func(ctx context.Context) (interface{}, error)) (interface{}, error) {
	if err := t.begin(); err != nil {
		return nil, err
	}
	ctx = context.WithValue(ctx, &txKey, t.db)
	v, err := f(ctx)
	if err != nil {
		return nil, t.rollback()
	}

	if err := t.commit(); err != nil {
		t.rollback()
		return nil, err
	}

	return v, nil
}

func (t *tx) begin() error {
	return t.db.Exec("SAVEPOINT SPTEST").Error
}

func (t *tx) rollback() error {
	return t.db.Exec("ROLLBACK TO SAVEPOINT SPTEST").Error
}

func (t *tx) commit() error {
	return t.db.Exec("RELEASE SAVEPOINT SPTEST").Error
}
