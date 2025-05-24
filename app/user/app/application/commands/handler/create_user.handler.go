package handler

import (
	"context"
	"user/app/application/commands/command"
	"user/app/domain/usecases"
)

type CreateUserCommandHandler struct {
	userUsecase usecases.UserUsecase
}

func NewCreateUserCommandHandler(userUsecase usecases.UserUsecase) *CreateUserCommandHandler {
	return &CreateUserCommandHandler{
		userUsecase: userUsecase,
	}
}

func (h *CreateUserCommandHandler) Handle(ctx context.Context, cmd *command.CreateUserCommand) (*command.CreateUserCommandResult, error) {
	return h.userUsecase.CreateUser(ctx, cmd)
}
