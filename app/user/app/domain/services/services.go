package services

import (
	"user/package/settings"
)

type Service interface {
	User() UserService
	Address() AddressService
}

type service struct {
	user    UserService
	address AddressService
}

func NewService(cfg *settings.Config) (Service, error) {
	return &service{
		user:    NewUserService(),
		address: NewAddressService(),
	}, nil
}

func (s *service) User() UserService {
	return s.user
}

func (s *service) Address() AddressService {
	return s.address
}
