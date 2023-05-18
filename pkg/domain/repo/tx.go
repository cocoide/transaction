package repository

import (
	"context"

	"github.com/cocoide/transaction/key"

	"gorm.io/gorm"
)

type TxRepo interface {
	DoInTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}

type txRepo struct {
	db *gorm.DB
}

func NewTxRepo(db *gorm.DB) TxRepo {
	return &txRepo{db: db}
}

func (r *txRepo) DoInTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	tx := r.db.Begin()
	// contextにkeyを元にtxを設定
	c := context.WithValue(ctx, key.TxKey, tx)
	var done bool

	defer func() {
		if !done {
			_ = tx.Rollback()
		}
	}()

	if err := fn(c); err != nil {
		return err
	}
	// ※タイムアウト時もロールバックする
	if ctx.Err() == context.DeadlineExceeded {
		return context.DeadlineExceeded
	}

	done = true
	return tx.Commit().Error
}
