package dto

import "errors"

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	AccessToken           string `json:"access_token"`
	RefreshToken          string `json:"refresh_token"`
	AccessTokenExpiresIn  int64  `json:"access_token_expires_in"`
	RefreshTokenExpiresIn int64  `json:"refresh_token_expires_in"`
}

func (l *LoginRequest) Validate() error {
	if l.Username == "" {
		return errors.New("username cannot be empty")
	}
	if l.Password == "" {
		return errors.New("password cannot be empty")
	}
	return nil
}

func (l *LoginRequest) GetUsername() string {
	return l.Username
}

func (l *LoginRequest) GetPassword() string {
	return l.Password
}
