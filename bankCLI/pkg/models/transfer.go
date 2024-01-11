package models

type Transfer struct {
	Sender    *Clients
	Recipient *Clients
	Amount    float64
}
