package handler

import (
	"context"
	"user/app/domain/dto"
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

func (h *LoginCommandHandler) Handle(ctx context.Context, cmd *dto.LoginRequest) (*dto.LoginResponse, error) {
	return h.userService.Login(ctx, cmd)
}
