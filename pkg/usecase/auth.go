package usecase

import (
	"context"

	"github.com/cocoide/transaction/pkg/domain/model"
	repo "github.com/cocoide/transaction/pkg/domain/repo"
)

type AuthUseCase interface {
	RegisterAndSendEmail(u *model.User, ctx context.Context) error
}
type authUseCase struct {
	ur repo.UserRepo
	tx repo.TxRepo
}

func NewAuthUseCase(ur repo.UserRepo, tx repo.TxRepo) AuthUseCase {
	return &authUseCase{ur: ur, tx: tx}
}

func (au *authUseCase) RegisterAndSendEmail(u *model.User, ctx context.Context) error {
	if err := au.tx.DoInTransaction(ctx, func(ctx context.Context) error {
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
