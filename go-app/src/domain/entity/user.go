package entity

import "time"

type User struct {
	ID              int        `json:"id"`
	Email           string     `json:"email"`
	Password        []byte     `json:"-"`
	EmailVerifiedAt *time.Time `json:"email_verified_at"`
}

func (u *User) SetEmailVerifiedAt() {
	now := time.Now()
	u.EmailVerifiedAt = &now
}
