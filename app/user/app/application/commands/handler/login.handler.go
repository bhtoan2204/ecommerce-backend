package handler

import (
	"context"
	"user/app/application/commands/command"
	"user/app/domain/services"
)

type LoginCommandHandler struct {
	userService services.UserService
}

func NewLoginCommandHandler(userService services.UserService) *LoginCommandHandler {
	return &LoginCommandHandler{
		userService: userService,
	}
}

func (h *LoginCommandHandler) Handle(ctx context.Context, cmd *command.LoginCommand) (*command.LoginCommandResult, error) {
	return h.userService.Login(ctx, cmd)
}
