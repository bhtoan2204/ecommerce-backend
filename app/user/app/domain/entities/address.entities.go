package entities

type Address struct {
	id        string
	userID    int64
	line1     string
	line2     string
	city      string
	state     string
	zipCode   string
	country   string
	isDefault bool
}

// Constructor
func NewAddress(userID int64, line1, line2, city, state, zipCode, country string, isDefault bool) *Address {
	return &Address{
		userID:    userID,
		line1:     line1,
		line2:     line2,
		city:      city,
		state:     state,
		zipCode:   zipCode,
		country:   country,
		isDefault: isDefault,
	}
}

// Getters
func (a *Address) ID() string {
	return a.id
}

func (a *Address) UserID() int64 {
	return a.userID
}

func (a *Address) Line1() string {
	return a.line1
}

func (a *Address) Line2() string {
	return a.line2
}

func (a *Address) City() string {
	return a.city
}

func (a *Address) State() string {
	return a.state
}

func (a *Address) ZipCode() string {
	return a.zipCode
}

func (a *Address) Country() string {
	return a.country
}

func (a *Address) IsDefault() bool {
	return a.isDefault
}

// Setters
func (a *Address) SetLine1(line1 string) {
	a.line1 = line1
}

func (a *Address) SetLine2(line2 string) {
	a.line2 = line2
}

func (a *Address) SetCity(city string) {
	a.city = city
}

func (a *Address) SetState(state string) {
	a.state = state
}

func (a *Address) SetZipCode(zip string) {
	a.zipCode = zip
}

func (a *Address) SetCountry(country string) {
	a.country = country
}

func (a *Address) SetIsDefault(isDefault bool) {
	a.isDefault = isDefault
}
