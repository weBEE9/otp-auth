package model

import "time"

type User struct {
	ID            int64     `db:"id"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
	Name          string    `db:"name"`
	Email         string    `db:"email"`
	PhoneNumber   string    `db:"phone_number"`
	PhoneVerified bool      `db:"phone_verified"`
}
