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

func (t *tx) DoInTx(ctx context.Context, f func(ctx context.Context) error) error {
	tx := t.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	ctx = context.WithValue(ctx, &txKey, tx)

	err := f(ctx)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func GetTx(ctx context.Context) (*gorm.DB, bool) {
	tx, ok := ctx.Value(&txKey).(*gorm.DB)
	return tx, ok
}
