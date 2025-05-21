package services

import (
	"context"
	"errors"
	"user/app/domain/entities"
	"user/app/infrastructure/persistent/postgresql/repository"
)

type AddressService interface {
	GetAddressByUserId(ctx context.Context, userId string) ([]*entities.Address, error)
	GetAddressById(ctx context.Context, id string) (*entities.Address, error)
	GetAllAddresses(ctx context.Context) ([]*entities.Address, error)
	CreateAddress(ctx context.Context, address *entities.Address) (*entities.Address, error)
	UpdateAddress(ctx context.Context, id string, address *entities.Address) (*entities.Address, error)
	DeleteAddress(ctx context.Context, id string) error
}

type addressService struct {
	addressRepository repository.AddressRepository
}

func NewAddressService(addressRepository repository.AddressRepository) AddressService {
	return &addressService{
		addressRepository: addressRepository,
	}
}

func (as *addressService) GetAddressByUserId(ctx context.Context, userId string) ([]*entities.Address, error) {
	return nil, errors.New("GetAddressByUserId: not implemented yet")
}

func (as *addressService) GetAddressById(ctx context.Context, id string) (*entities.Address, error) {
	return nil, errors.New("GetAddressById: not implemented yet")
}

func (as *addressService) GetAllAddresses(ctx context.Context) ([]*entities.Address, error) {
	return nil, errors.New("GetAllAddresses: not implemented yet")
}

func (as *addressService) CreateAddress(ctx context.Context, address *entities.Address) (*entities.Address, error) {
	return nil, errors.New("CreateAddress: not implemented yet")
}

func (as *addressService) UpdateAddress(ctx context.Context, id string, address *entities.Address) (*entities.Address, error) {
	return nil, errors.New("UpdateAddress: not implemented yet")
}

func (as *addressService) DeleteAddress(ctx context.Context, id string) error {
	return errors.New("DeleteAddress: not implemented yet")
}
