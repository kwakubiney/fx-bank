package models

import "time"

type User struct {
	ID           string    `json:"id" gorm:"default:gen_random_uuid()"`
	Username     string    `json:"username"`
	Password     string    `json:"password"`
	CreatedAt    time.Time `json:"created_at" sql:"type:timestamp without time zone"`
	LastModified time.Time `json:"last_modified" sql:"type:timestamp without time zone"`
}
