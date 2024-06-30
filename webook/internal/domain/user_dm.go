package domain

import "time"

type User struct {
	Id       int64     `json:"id"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
	Password string    `json:"password"`
	Ctime    time.Time `json:"ctime"`
}
