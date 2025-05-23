package command

import (
	"time"
	"user/app/presentation/dto"
)

type CreateUserCommand struct {
	Email       string    `json:"email" validate:"required,email"`
	Password    string    `json:"password" validate:"required,min=6"`
	FirstName   string    `json:"first_name" validate:"required"`
	LastName    string    `json:"last_name" validate:"required"`
	PhoneNumber string    `json:"phone_number" validate:"omitempty"`
	Avatar      string    `json:"avatar" validate:"omitempty,url"`
	BirthDate   time.Time `json:"birth_date" validate:"omitempty"`
}

func (r *CreateUserCommand) CommandName() string {
	return "CreateUserCommand"
}

type CreateUserCommandResult struct {
	Message string
}

func (r *CreateUserCommandResult) ToDto() *dto.CreateUserResponse {
	return &dto.CreateUserResponse{
		Message: r.Message,
	}
}
