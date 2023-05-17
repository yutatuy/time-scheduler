package repository

import (
	"context"
	"go-app/src/domain/entity"
	gormpkg "go-app/src/infrastructure/gorm"
	"go-app/src/infrastructure/gorm/model"

	"gorm.io/gorm"
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}

type UserRepository interface {
	FindByID(id int) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	Create(ctx context.Context, user entity.User) (*entity.User, error)
	Update(user *entity.User) (*entity.User, error)
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func (r *UserRepositoryImpl) FindByID(id int) (*entity.User, error) {
	var gormUser model.GormUser
	if err := r.db.Find(&gormUser, id).Error; err != nil {
		return nil, err
	}

	return createFromGormUser(&gormUser), nil
}

func (r *UserRepositoryImpl) FindByEmail(email string) (*entity.User, error) {
	var gormUser model.GormUser
	if err := r.db.Where("email = ?", email).First(&gormUser).Error; err != nil {
		return nil, err
	}

	return createFromGormUser(&gormUser), nil
}

func (r *UserRepositoryImpl) Create(ctx context.Context, user entity.User) (*entity.User, error) {
	db, ok := gormpkg.GetTx(ctx)
	if !ok {
		db = r.db
	}

	gormUser := &model.GormUser{
		Email:           user.Email,
		Password:        []byte(user.Password),
		EmailVerifiedAt: user.EmailVerifiedAt,
	}

	result := db.Create(gormUser)
	if result.Error != nil {
		return nil, result.Error
	}

	return createFromGormUser(gormUser), nil
}

func (r *UserRepositoryImpl) Update(user *entity.User) (*entity.User, error) {
	var gormUser model.GormUser
	if err := r.db.Find(&gormUser, user.ID).Error; err != nil {
		return nil, err
	}

	result := r.db.Model(&gormUser).Updates(
		model.GormUser{
			Email:           user.Email,
			EmailVerifiedAt: user.EmailVerifiedAt,
		})
	if result.Error != nil {
		return nil, result.Error
	}

	return createFromGormUser(&gormUser), nil
}

func createFromGormUser(u *model.GormUser) *entity.User {
	return &entity.User{
		ID:       u.ID,
		Email:    u.Email,
		Password: u.Password,
	}
}
