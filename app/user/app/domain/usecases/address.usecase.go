package usecases

import (
	"context"
	"errors"
	"user/app/domain/entities"
	"user/app/infrastructure/persistent/postgresql/repository"
)

type AddressUsecase interface {
	GetAddressByUserId(ctx context.Context, userId string) ([]*entities.Address, error)
	GetAddressById(ctx context.Context, id string) (*entities.Address, error)
	GetAllAddresses(ctx context.Context) ([]*entities.Address, error)
	CreateAddress(ctx context.Context, address *entities.Address) (*entities.Address, error)
	UpdateAddress(ctx context.Context, id string, address *entities.Address) (*entities.Address, error)
	DeleteAddress(ctx context.Context, id string) error
}

type addressUsecase struct {
	addressRepository repository.AddressRepository
}

func NewAddressUsecase(addressRepository repository.AddressRepository) AddressUsecase {
	return &addressUsecase{
		addressRepository: addressRepository,
	}
}

func (as *addressUsecase) GetAddressByUserId(ctx context.Context, userId string) ([]*entities.Address, error) {
	return nil, errors.New("GetAddressByUserId: not implemented yet")
}

func (as *addressUsecase) GetAddressById(ctx context.Context, id string) (*entities.Address, error) {
	return nil, errors.New("GetAddressById: not implemented yet")
}

func (as *addressUsecase) GetAllAddresses(ctx context.Context) ([]*entities.Address, error) {
	return nil, errors.New("GetAllAddresses: not implemented yet")
}

func (as *addressUsecase) CreateAddress(ctx context.Context, address *entities.Address) (*entities.Address, error) {
	return nil, errors.New("CreateAddress: not implemented yet")
}

func (as *addressUsecase) UpdateAddress(ctx context.Context, id string, address *entities.Address) (*entities.Address, error) {
	return nil, errors.New("UpdateAddress: not implemented yet")
}

func (as *addressUsecase) DeleteAddress(ctx context.Context, id string) error {
	return errors.New("DeleteAddress: not implemented yet")
}
