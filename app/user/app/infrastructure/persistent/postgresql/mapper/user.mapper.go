package mapper

import (
	"user/app/domain/entities"
	"user/app/infrastructure/persistent/postgresql/model"
)

func UserToModel(user *entities.User) *model.User {
	addresses := make([]model.Address, len(user.Address()))
	for i, address := range user.Address() {
		addresses[i] = *AddressToModel(address)
	}
	return &model.User{
		Email:       user.Email(),
		Password:    user.PasswordHash(),
		FirstName:   user.FirstName(),
		LastName:    user.LastName(),
		PhoneNumber: user.PhoneNumber(),
		Addresses:   addresses,
	}
}

func UserToEntity(user *model.User) *entities.User {
	addresses := make([]*entities.Address, len(user.Addresses))
	for i, address := range user.Addresses {
		addresses[i] = AddressToEntity(&address)
	}
	return entities.NewUserWithID(
		user.ID,
		user.Email,
		user.Password,
		user.FirstName,
		user.LastName,
		user.PhoneNumber,
		user.BirthDate,
		addresses,
		user.Avatar,
		user.PinCode,
		user.Role,
	)
}
