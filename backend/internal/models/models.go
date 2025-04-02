package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Username  string    `json:"username"`
	Role      string    `json:"role"`
	AvatarURL string    `json:"avatar_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUser(ID uuid.UUID, fullName, email, password, username, role, avatarURL string) *User {
	now := time.Now()
	return &User{
		ID:        ID,
		FullName:  fullName,
		Email:     email,
		Password:  password,
		Username:  username,
		AvatarURL: avatarURL,
		Role:      role,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
