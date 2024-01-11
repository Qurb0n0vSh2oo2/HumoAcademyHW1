package models

type Clients struct {
	Name        string
	Age         int
	PhoneNumber string
	Balance     float64
	City        *City
}
