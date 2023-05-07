package repository

import (
	entity "go-app/src/domain/entity/user"
	"go-app/src/infrastructure/gorm/connection"
	"go-app/src/infrastructure/gorm/model"
)

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

type UserRepository interface {
	CreateByEmail(user entity.User) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
}

type UserRepositoryImpl struct {
}

func (r *UserRepositoryImpl) CreateByEmail(user entity.User) (*entity.User, error) {
	db := connection.DBConnect()
	gormUser := &model.GormUser{
		Email:    user.Email,
		Password: []byte(user.Password),
	}

	result := db.Create(gormUser)
	if result.Error != nil {
		return nil, result.Error
	}

	return createFromGorm(gormUser), nil
}

func (r *UserRepositoryImpl) FindByEmail(email string) (*entity.User, error) {
	db := connection.DBConnect()
	var user model.GormUser
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return createFromGorm(&user), nil
}

func createFromGorm(u *model.GormUser) *entity.User {
	return &entity.User{
		Email:    u.Email,
		Password: u.Password,
	}
}
