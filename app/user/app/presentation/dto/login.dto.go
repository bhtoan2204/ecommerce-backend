package dto

import (
	"errors"
	"user/proto/user"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (r *LoginRequest) CommandName() string {
	return "LoginCommand"
}

func (l *LoginRequest) Validate() error {
	if l.Email == "" {
		return errors.New("email cannot be empty")
	}
	if l.Password == "" {
		return errors.New("password cannot be empty")
	}
	return nil
}

type LoginResponse struct {
	AccessToken           string `json:"access_token"`
	RefreshToken          string `json:"refresh_token"`
	AccessTokenExpiresIn  int64  `json:"access_token_expires_in"`
	RefreshTokenExpiresIn int64  `json:"refresh_token_expires_in"`
}

func (res *LoginResponse) ToPb() *user.LoginResponse {
	return &user.LoginResponse{
		AccessToken:           res.AccessToken,
		RefreshToken:          res.RefreshToken,
		AccessTokenExpiresIn:  res.AccessTokenExpiresIn,
		RefreshTokenExpiresIn: res.RefreshTokenExpiresIn,
	}
}
