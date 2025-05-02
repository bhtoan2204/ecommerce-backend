package interfaces

import "user/app/domain/entities"

type AddressService interface {
	GetAddressByUserId(userId string) ([]*entities.Address, error)
	GetAddressById(id string) (*entities.Address, error)
	GetAllAddresses() ([]*entities.Address, error)
	CreateAddress(address *entities.Address) (*entities.Address, error)
	UpdateAddress(id string, address *entities.Address) (*entities.Address, error)
	DeleteAddress(id string) error
}
