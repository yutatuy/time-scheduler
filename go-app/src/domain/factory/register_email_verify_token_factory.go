package factory

import (
	"go-app/src/domain/entity"
	"go-app/src/domain/repository"
	"time"
)

func NewRegisterEmailVerifyToken(userID int, email string) entity.RegisterEmailVerifyToken {
	tokenCreator := repository.NewTokenCreator()
	token := tokenCreator.Create()
	return entity.RegisterEmailVerifyToken{
		UserID:    userID,
		Email:     email,
		Token:     token,
		ExpiredAt: time.Now().Add(time.Hour * 24),
	}
}
