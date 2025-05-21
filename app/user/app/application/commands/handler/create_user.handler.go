package handler

import (
	"context"
	"user/app/domain/dto"
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

func (h *CreateUserCommandHandler) Handle(ctx context.Context, cmd *dto.CreateUserRequest) (*dto.CreateUserResponse, error) {
	return h.userService.CreateUser(ctx, cmd)
}
