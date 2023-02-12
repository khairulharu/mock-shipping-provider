package primitive

import "regexp"

type Address struct {
	Name        string     `json:"name"`
	PhoneNumber string     `json:"phone_number"`
	Address     string     `json:"address"`
	City        string     `json:"city"`
	State       string     `json:"state"`
	Country     string     `json:"country"`
	PostalCode  string     `json:"postal_code"`
	Coordinate  Coordinate `json:"coordinate"`
}

var AddressNamePattern = regexp.MustCompile(`[\w\s]+`)
var AddressPhoneNumberPattern = regexp.MustCompile(`^\+\d+$`)
