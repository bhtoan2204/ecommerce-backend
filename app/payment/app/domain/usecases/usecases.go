package usecases

import (
	"payment/app/infrastructure/persistent/postgresql/repository"
	"payment/package/settings"
)

type Usecase interface {
}

type usecase struct {
}

func NewUsecase(cfg *settings.Config, repository *repository.Repository) (Usecase, error) {
	return &usecase{}, nil
}
