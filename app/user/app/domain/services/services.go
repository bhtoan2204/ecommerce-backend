package services

import (
	"user/app/infrastructure/persistent/postgresql/repository"
	"user/package/jwt_utils"
	"user/package/settings"
)

type Service interface {
	UserService() UserService
	AddressService() AddressService
}

type service struct {
	userService    UserService
	addressService AddressService
}

func NewService(cfg *settings.Config, repository *repository.Repository) (Service, error) {
	userService := NewUserService(
		(*repository).UserRepository(),
		*jwt_utils.NewJWTUtils(
			cfg.SecurityConfig.JWTAccessSecret,
			cfg.SecurityConfig.JWTRefreshSecret,
			int32(cfg.SecurityConfig.JWTAccessExpiration),
			int32(cfg.SecurityConfig.JWTRefreshExpiration)))

	addressService := NewAddressService((*repository).AddressRepository())
	return &service{
		userService:    userService,
		addressService: addressService,
	}, nil
}

func (s *service) UserService() UserService {
	return s.userService
}

func (s *service) AddressService() AddressService {
	return s.addressService
}
