package handler

import (
	"context"
	"user/app/application/commands/command"
	"user/app/domain/services"
)

type CreateUserCommandHandler struct {
	userService services.UserService
}

func NewCreateUserCommandHandler(userService services.UserService) *CreateUserCommandHandler {
	return &CreateUserCommandHandler{
		userService: userService,
	}
}

func (h *CreateUserCommandHandler) Handle(ctx context.Context, cmd *command.CreateUserCommand) (*command.CreateUserCommandResult, error) {
	return h.userService.CreateUser(ctx, cmd)
}
