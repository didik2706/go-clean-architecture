package entity

import "time"

type User struct {
	Id        int
	Name      string
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
