package interfaces

import "user/app/domain/dto"

type UserService interface {
	Login(request *dto.LoginRequest) (*dto.LoginResponse, error)
}
