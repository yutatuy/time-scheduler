package repository

import (
	"context"
	"go-app/src/domain/entity"
	gormpkg "go-app/src/infrastructure/gorm"
	"go-app/src/infrastructure/gorm/model"

	"gorm.io/gorm"
)

func NewRegisterEmailVerifyTokenRepository(db *gorm.DB) RegisterEmailVerifyTokenRepository {
	return &RegisterEmailVerifyTokenRepositoryImpl{
		db: db,
	}
}

type RegisterEmailVerifyTokenRepository interface {
	Create(ctx context.Context, e entity.RegisterEmailVerifyToken) (*entity.RegisterEmailVerifyToken, error)
	FindByToken(t string) (*entity.RegisterEmailVerifyToken, error)
}

type RegisterEmailVerifyTokenRepositoryImpl struct {
	db *gorm.DB
}

func (r *RegisterEmailVerifyTokenRepositoryImpl) Create(ctx context.Context, e entity.RegisterEmailVerifyToken) (*entity.RegisterEmailVerifyToken, error) {
	db, ok := gormpkg.GetTx(ctx)
	if !ok {
		db = r.db
	}

	gormRegisterEmailVerifyToken := &model.GormRegisterEmailVerifyToken{
		ID:        e.ID,
		UserID:    e.UserID,
		Email:     e.Email,
		Token:     e.Token,
		ExpiredAt: e.ExpiredAt,
	}

	result := db.Create(gormRegisterEmailVerifyToken)
	if result.Error != nil {
		return nil, result.Error
	}

	return createFromGormRegisterEmailVerifyToken(gormRegisterEmailVerifyToken), nil
}

func (r *RegisterEmailVerifyTokenRepositoryImpl) FindByToken(t string) (*entity.RegisterEmailVerifyToken, error) {
	var token model.GormRegisterEmailVerifyToken

	if err := r.db.Model(&model.GormRegisterEmailVerifyToken{}).Where("token = ?", t).First(&token).Error; err != nil {
		return nil, err
	}

	return createFromGormRegisterEmailVerifyToken(&token), nil
}

func createFromGormRegisterEmailVerifyToken(g *model.GormRegisterEmailVerifyToken) *entity.RegisterEmailVerifyToken {
	return &entity.RegisterEmailVerifyToken{
		ID:        g.ID,
		UserID:    g.UserID,
		Email:     g.Email,
		Token:     g.Token,
		ExpiredAt: g.ExpiredAt,
	}
}
