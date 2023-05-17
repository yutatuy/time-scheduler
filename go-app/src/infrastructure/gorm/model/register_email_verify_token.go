package model

import (
	"time"

	"gorm.io/gorm"
)

type GormRegisterEmailVerifyToken struct {
	gorm.Model
	ID        int       `gorm:"primaryKey"`
	UserID    int       `gorm:"not null"`
	Email     string    `gorm:"not null"`
	Token     string    `gorm:"not null;unique"`
	ExpiredAt time.Time `gorm:"not null"`
}

func (GormRegisterEmailVerifyToken) TableName() string {
	return "register_email_verify_tokens"
}
