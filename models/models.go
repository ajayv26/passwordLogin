package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type User struct {
	ID           int64     `json:"id"`
	Code         string    `json:"code"`
	FirstName    string    `json:"firstName"`
	LastName     string    `json:"lastName"`
	Email        string    `json:"email"`
	Phone        string    `json:"phone"`
	PasswordHash string    `json:"passwordHash"`
	IsAdmin      bool      `json:"isAdmin"`
	IsArchived   bool      `json:"isArchived"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type UserReq struct {
	ID        int64     `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Auther struct {
	ID    int64     `json:"id"`
	Name  string    `json:"name"`
	Token uuid.UUID `json:"token"`
}

type AuthToken struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"userID"`
	Token     uuid.UUID `json:"token"`
	ExpiresAt time.Time `json:"expiresAt"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
