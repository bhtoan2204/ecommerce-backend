package dto

import (
	"errors"
	"time"
	"user/proto/user"
)

type CreateUserRequest struct {
	Email       string    `json:"email" validate:"required,email"`
	Password    string    `json:"password" validate:"required,min=6"`
	FirstName   string    `json:"first_name" validate:"required"`
	LastName    string    `json:"last_name" validate:"required"`
	PhoneNumber string    `json:"phone_number" validate:"omitempty"`
	Avatar      string    `json:"avatar" validate:"omitempty,url"`
	BirthDate   time.Time `json:"birth_date" validate:"omitempty"`
}

func (r *CreateUserRequest) CommandName() string {
	return "CreateUserRequest"
}

func (l *CreateUserRequest) Validate() error {
	if l.Email == "" {
		return errors.New("email cannot be empty")
	}
	if l.Password == "" {
		return errors.New("password cannot be empty")
	}
	return nil
}

type CreateUserResponse struct {
	Message string
}

func (r *CreateUserResponse) ToPb() *user.CreateUserResponse {
	return &user.CreateUserResponse{
		Message: r.Message,
	}
}
