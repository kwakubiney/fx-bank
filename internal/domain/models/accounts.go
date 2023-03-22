package models

import (
	"time"
)

type Account struct {
	ID           string    `json:"id" gorm:"default:gen_random_uuid()"`
	Name         string    `json:"name"`
	Balance      int64     `json:"balance" sql:"type:decimal(20,8);"`
	CreatedAt    time.Time `json:"created_at" sql:"type:timestamp without time zone"`
	LastModified time.Time `json:"last_modified" sql:"type:timestamp without time zone"`
	UserID       string    `json:"user_id" sql:""`
}
