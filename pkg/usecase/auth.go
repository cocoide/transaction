package usecase

import (
	"context"
	"time"

	"github.com/cocoide/transaction/pkg/domain/model"
	repository "github.com/cocoide/transaction/pkg/domain/repo"
)

type AuthUseCase interface {
	RegisterAndSendEmail(u *model.User, ctx context.Context) error
}
type authUseCase struct {
	ur repository.UserRepo
	tx repository.TxRepo
}

func NewAuthUseCase(ur repository.UserRepo, tx repository.TxRepo) AuthUseCase {
	return &authUseCase{ur: ur, tx: tx}
}

func (au *authUseCase) RegisterAndSendEmail(u *model.User, ctx context.Context) error {
	if err := au.tx.DoInTransaction(ctx, func(ctx context.Context) error {
		time.Sleep(3 * time.Second)
		if err := au.ur.Create(ctx, u); err != nil {
			return err
		}
		// メール送信の処理をトランザクション中に実行
		// if err:= eg.SendRegisterComfirmEmail; err != nil {
		// 	return err
		// }
		return nil
	}); err != nil {
		return err
	}
	return nil
}
