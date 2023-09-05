package models

import "time"

type User struct {
	Id        int       `db:"id" json:"id,omitempty"`
	Username  string    `db:"username" json:"username,omitempty"`
	Password  string    `db:"password" json:"password,omitempty"`
	Salt      string    `db:"salt" json:"salt,omitempty"`
	Email     string    `db:"email" json:"email,omitempty"`
	CreatedAt time.Time `db:"created_at" json:"createdAt,omitempty"`
}
