package handler

import (
	"context"
	"time"

	"github.com/cocoide/transaction/pkg/domain/model"
	"github.com/cocoide/transaction/pkg/usecase"

	"github.com/labstack/echo/v4"
)

type AuthHandler interface {
	Register(c echo.Context) error
}
type authHandler struct {
	au usecase.AuthUseCase
}

func NewAuthHandler(au usecase.AuthUseCase) AuthHandler {
	return &authHandler{au: au}
}

func (h *authHandler) Register(c echo.Context) error {
	// Timeout within 2 Seconds
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	// Bind with Echo
	var u model.User
	if err := c.Bind(&u); err != nil {
		return c.JSON(400, err.Error())
	}
	if err := h.au.RegisterAndSendEmail(&u, ctx); err != nil {
		return c.JSON(500, err.Error())
	}
	return c.JSON(200, "register successfully completed")
}
