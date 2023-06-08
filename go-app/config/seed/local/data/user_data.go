package seed_local_data

import (
	"fmt"
	"go-app/src/infrastructure/gorm/model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	users := []*model.GormUser{}
	for i := 1; i <= 10; i++ {
		email := fmt.Sprintf("user%d@info.com", i)
		user := &model.GormUser{
			Email:    email,
			Password: hashedPassword,
		}
		users = append(users, user)
	}

	if err := db.Create(&users).Error; err != nil {
		fmt.Printf("%+v", err)
	}

	return nil
}
