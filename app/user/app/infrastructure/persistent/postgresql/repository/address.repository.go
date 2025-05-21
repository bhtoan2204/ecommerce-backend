package repository

import (
	"context"
	"user/app/infrastructure/persistent/postgresql/model"

	"gorm.io/gorm"
)

var _ AddressRepository = (*addressRepository)(nil)

type AddressRepository interface {
	Create(ctx context.Context, address *model.Address) (*model.Address, error)
	Update(ctx context.Context, address *model.Address) (*model.Address, error)
	Delete(ctx context.Context, address *model.Address) error
}

type addressRepository struct {
	db *gorm.DB
}

func newAddressRepository(db *gorm.DB) AddressRepository {
	return &addressRepository{
		db: db,
	}
}

func (r *addressRepository) Create(ctx context.Context, address *model.Address) (*model.Address, error) {
	panic("unimplemented")
}

func (r *addressRepository) Update(ctx context.Context, address *model.Address) (*model.Address, error) {
	panic("unimplemented")
}

func (r *addressRepository) Delete(ctx context.Context, address *model.Address) error {
	panic("unimplemented")
}
