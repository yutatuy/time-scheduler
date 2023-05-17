package gormpkg

import (
	"context"
	"go-app/src/domain/shared"

	"gorm.io/gorm"
)

type tx struct {
	db *gorm.DB
}

func NewTransaction(db *gorm.DB) shared.Transaction {
	return &tx{db: db}
}

var txKey = struct{}{}

func (t *tx) DoInTx(ctx context.Context, f func(ctx context.Context) (interface{}, error)) (interface{}, error) {
	tx := t.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	ctx = context.WithValue(ctx, &txKey, tx)

	v, err := f(ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	return v, nil
}

func GetTx(ctx context.Context) (*gorm.DB, bool) {

	tx, ok := ctx.Value(&txKey).(*gorm.DB)
	return tx, ok
}
