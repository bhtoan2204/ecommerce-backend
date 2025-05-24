package usecases

import (
	"user/app/infrastructure/persistent/postgresql/repository"
	"user/package/jwt_utils"
	"user/package/settings"
)

type Usecase interface {
	UserUsecase() UserUsecase
	AddressUsecase() AddressUsecase
}

type usecase struct {
	userUsecase    UserUsecase
	addressUsecase AddressUsecase
}

func NewUsecase(cfg *settings.Config, repository *repository.Repository) (Usecase, error) {
	userUsecase := NewUserUsecase(
		(*repository).UserRepository(),
		*jwt_utils.NewJWTUtils(
			cfg.SecurityConfig.JWTAccessSecret,
			cfg.SecurityConfig.JWTRefreshSecret,
			int32(cfg.SecurityConfig.JWTAccessExpiration),
			int32(cfg.SecurityConfig.JWTRefreshExpiration)))

	addressUsecase := NewAddressUsecase((*repository).AddressRepository())
	return &usecase{
		userUsecase:    userUsecase,
		addressUsecase: addressUsecase,
	}, nil
}

func (s *usecase) UserUsecase() UserUsecase {
	return s.userUsecase
}

func (s *usecase) AddressUsecase() AddressUsecase {
	return s.addressUsecase
}
