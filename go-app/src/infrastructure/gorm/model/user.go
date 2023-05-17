package model

import (
	"time"

	"gorm.io/gorm"
)

type GormUser struct {
	gorm.Model
	ID              int    `gorm:"not null;unique"`
	Email           string `gorm:"not null;unique"`
	Password        []byte `gorm:"not null"`
	EmailVerifiedAt *time.Time
}

func (GormUser) TableName() string {
	return "users"
}
