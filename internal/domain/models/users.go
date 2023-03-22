package models

import "time"

type User struct {
	ID          string    `json:"id" gorm:"default:gen_random_uuid()"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Email       string    `json:"email"`
	DateCreated time.Time `json:"dateCreated" sql:"type:timestamp without time zone"`
}
