package models

import "time"

type Transaction struct {
	ID                string    `json:"id" gorm:"default:gen_random_uuid()"`
	Credit            string    `json:"credit"`
	Debit             string    `json:"debit"`
	Amount            int64     `json:"amount"`
	SenderAccountID   string    `json:"sender_account_id"`
	ReceiverAccountID string    `json:"receiver_account_id"`
	CreatedAt         time.Time `json:"created_at" sql:"type:timestamp without time zone"`
}
