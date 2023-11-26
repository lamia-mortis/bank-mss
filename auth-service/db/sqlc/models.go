// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package db

import (
	"time"
)

type User struct {
	Name              string    `json:"name"`
	Surname           string    `json:"surname"`
	Username          string    `json:"username"`
	Password          string    `json:"password"`
	Email             string    `json:"email"`
	IsEmailVerified   bool      `json:"is_email_verified"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}