package mapper

import (
	"user/app/domain/entities"
	"user/app/infrastructure/persistent/postgresql/model"
)

func AddressToModel(address *entities.Address) *model.Address {
	return &model.Address{
		UserID:    address.UserID(),
		Line1:     address.Line1(),
		Line2:     address.Line2(),
		City:      address.City(),
		State:     address.State(),
		ZipCode:   address.ZipCode(),
		Country:   address.Country(),
		IsDefault: address.IsDefault(),
	}
}

func AddressToEntity(address *model.Address) *entities.Address {
	return entities.NewAddress(
		address.UserID,
		address.Line1,
		address.Line2,
		address.City,
		address.State,
		address.ZipCode,
		address.Country,
		address.IsDefault,
	)
}
