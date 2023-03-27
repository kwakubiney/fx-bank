package models

import (
	"errors"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"time"
)

var (
	ErrInsufficientBalance = errors.New("insufficient balance in account")
)

type Account struct {
	ID           string          `json:"id" gorm:"default:gen_random_uuid()"`
	Name         string          `json:"name"`
	Balance      decimal.Decimal `json:"balance" sql:"type:decimal(20,8);"`
	Currency     Currency        `json:"currency"`
	CreatedAt    time.Time       `json:"created_at" sql:"type:timestamp without time zone"`
	LastModified time.Time       `json:"last_modified" sql:"type:timestamp without time zone"`
	UserID       uuid.UUID       `gorm:"foreignKey:user_id" json:"user_id"`
}

func (a *Account) Deposit(amount decimal.Decimal) decimal.Decimal {
	return a.Balance.Add(amount)
}

func (a *Account) Withdraw(amount decimal.Decimal, rate decimal.Decimal) error {
	if a.Balance.LessThan(amount.Mul(rate)) {
		return ErrInsufficientBalance
	}

	//log.Println(fmt.Sprintf("rate x balance = %d"), amount*rate)
	a.Balance = a.Balance.Sub(rate.Mul(amount))
	return nil
}
