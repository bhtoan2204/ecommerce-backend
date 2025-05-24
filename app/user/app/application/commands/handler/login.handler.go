package handler

import (
	"context"
	"user/app/application/commands/command"
	"user/app/domain/usecases"
)

type LoginCommandHandler struct {
	userUsecase usecases.UserUsecase
}

func NewLoginCommandHandler(userUsecase usecases.UserUsecase) *LoginCommandHandler {
	return &LoginCommandHandler{
		userUsecase: userUsecase,
	}
}

func (h *LoginCommandHandler) Handle(ctx context.Context, cmd *command.LoginCommand) (*command.LoginCommandResult, error) {
	return h.userUsecase.Login(ctx, cmd)
}
