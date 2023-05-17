package entity

import (
	"time"
)

type RegisterEmailVerifyToken struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Email     string    `json:"email"`
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expired_at"`
}

func (r *RegisterEmailVerifyToken) CheckExpired() bool {
	return time.Now().After(r.ExpiredAt)
}
