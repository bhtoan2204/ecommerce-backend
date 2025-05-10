package interfaces

type Service interface {
	User() UserService
	Address() AddressService
}
