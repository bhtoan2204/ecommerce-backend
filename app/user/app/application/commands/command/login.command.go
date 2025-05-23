package command

import "user/app/presentation/dto"

type LoginCommand struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (r *LoginCommand) CommandName() string {
	return "LoginCommand"
}

type LoginCommandResult struct {
	AccessToken           string `json:"access_token"`
	RefreshToken          string `json:"refresh_token"`
	AccessTokenExpiresIn  int64  `json:"access_token_expires_in"`
	RefreshTokenExpiresIn int64  `json:"refresh_token_expires_in"`
}

func (res *LoginCommandResult) ToDto() *dto.LoginResponse {
	return &dto.LoginResponse{
		AccessToken:           res.AccessToken,
		RefreshToken:          res.RefreshToken,
		AccessTokenExpiresIn:  res.AccessTokenExpiresIn,
		RefreshTokenExpiresIn: res.RefreshTokenExpiresIn,
	}
}
