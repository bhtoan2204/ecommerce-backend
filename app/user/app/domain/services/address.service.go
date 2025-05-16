package services

import (
	"errors"
	"user/app/domain/entities"
)

type AddressService interface {
	GetAddressByUserId(userId string) ([]*entities.Address, error)
	GetAddressById(id string) (*entities.Address, error)
	GetAllAddresses() ([]*entities.Address, error)
	CreateAddress(address *entities.Address) (*entities.Address, error)
	UpdateAddress(id string, address *entities.Address) (*entities.Address, error)
	DeleteAddress(id string) error
}

type addressService struct{}

func NewAddressService() AddressService {
	return &addressService{}
}

func (as *addressService) GetAddressByUserId(userId string) ([]*entities.Address, error) {
	return nil, errors.New("GetAddressByUserId: not implemented yet")
}

func (as *addressService) GetAddressById(id string) (*entities.Address, error) {
	return nil, errors.New("GetAddressById: not implemented yet")
}

func (as *addressService) GetAllAddresses() ([]*entities.Address, error) {
	return nil, errors.New("GetAllAddresses: not implemented yet")
}

func (as *addressService) CreateAddress(address *entities.Address) (*entities.Address, error) {
	return nil, errors.New("CreateAddress: not implemented yet")
}

func (as *addressService) UpdateAddress(id string, address *entities.Address) (*entities.Address, error) {
	return nil, errors.New("UpdateAddress: not implemented yet")
}

func (as *addressService) DeleteAddress(id string) error {
	return errors.New("DeleteAddress: not implemented yet")
}
