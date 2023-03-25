package models

import "time"

type Status string

const (
	Pending   Status = "PENDING"
	Completed Status = "COMPLETED"
	Failed    Status = "FAILED"
)

type Transaction struct {
	ID                  string    `json:"id" gorm:"default:gen_random_uuid()"`
	Credit              int64     `json:"credit"`
	Debit               int64     `json:"debit"`
	ProviderName        string    `json:"providerName"`
	SenderAccountName   string    `json:"sender_account_name"`
	ReceiverAccountName string    `json:"receiver_account_name"`
	SenderCurrency      Currency  `json:"sender_currency"`
	ReceiverCurrency    Currency  `json:"receiver_currency"`
	CreatedAt           time.Time `json:"created_at" sql:"type:timestamp without time zone"`
	UpdatedAt           time.Time `json:"updated-at" sql:"type:timestamp without time zone"`
	Status              Status    `json:"status"`
	Rate                int64     `json:"rate"`
}
