package users

import "time"

type Users struct {
	ID             int
	Name           string
	Occupation     string
	PasswordHash   string
	Email          string
	Role           string
	AvatarFileName string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
