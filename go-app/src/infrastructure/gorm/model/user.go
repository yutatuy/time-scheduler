package model

import "gorm.io/gorm"

type GormUser struct {
	gorm.Model
	Email    string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
}

func (GormUser) TableName() string {
	return "users"
}
