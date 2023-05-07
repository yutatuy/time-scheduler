package model

import "gorm.io/gorm"

type GormUser struct {
	gorm.Model
	Id       int    `gorm:"not null;unique"`
	Email    string `gorm:"not null;unique"`
	Password []byte `gorm:"not null"`
}

func (GormUser) TableName() string {
	return "users"
}
