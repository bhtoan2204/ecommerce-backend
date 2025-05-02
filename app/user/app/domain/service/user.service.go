package service

import (
	"user/app/domain/dto"
	"user/app/domain/interfaces"
)

type userService struct {
}

func NewUserService() interfaces.UserService {
	return &userService{}
}
func (s *userService) Login(request *dto.LoginRequest) (*dto.LoginResponse, error) {
	// Implement the login logic here
	// For example, validate the user credentials and generate tokens
	// Return the LoginResponse with the generated tokens
	return &dto.LoginResponse{
		AccessToken:           "access_token",
		RefreshToken:          "refresh_token",
		AccessTokenExpiresIn:  3600,
		RefreshTokenExpiresIn: 7200,
	}, nil
}
