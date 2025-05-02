package service

import (
	"errors"
	"user/app/domain/entities"
	"user/app/domain/interfaces"
)

type addressService struct{}

func NewAddressService() interfaces.AddressService {
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
