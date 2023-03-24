package models

import "fmt"

type Currency string

const (
	UnitedStatesDollar Currency = "USD"
	Euro               Currency = "Euro"
	Naira              Currency = "Naira"
	Cedi               Currency = "GHC"
)

func (c Currency) toString() string {
	return fmt.Sprintf("%s", c)
}
