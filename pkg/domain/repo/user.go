package repository

import (
	"context"

	"github.com/cocoide/transaction/key"
	"github.com/cocoide/transaction/pkg/domain/model"

	"gorm.io/gorm"
)

type UserRepo interface {
	Create(ctx context.Context, u *model.User) error
}
type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{db: db}
}

func (ur *userRepo) Create(ctx context.Context, u *model.User) error {
	if tx, ok := ctx.Value(key.TxKey).(*gorm.DB); ok {
		return tx.Create(u).Error
	}
	return ur.db.Create(u).Error
}
