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
}

type UserRepositoryImpl struct {
}

func (r *UserRepositoryImpl) CreateByEmail(user entity.User) (*entity.User, error) {
	db := connection.DBConnect()
	gormUser := &model.GormUser{
		Email:    user.Email,
		Password: user.Password,
	}

	result := db.Create(gormUser)
	if result.Error != nil {
		return nil, result.Error
	}

	entity := &entity.User{
		Email:    gormUser.Email,
		Password: gormUser.Password,
	}
	return entity, nil
}
