package controller

import (
	"context"
	"go-app/src/adapter/convertor"
	"go-app/src/adapter/presenter"
	"go-app/src/application/usecase"
	"go-app/src/domain/repository"
	gormpkg "go-app/src/infrastructure/gorm"
	"go-app/src/infrastructure/gorm/connection"

	"github.com/gin-gonic/gin"
)

func RegisterByEmail(c *gin.Context) {
	db := connection.DBConnect()
	ctx := context.Background()
	tx := gormpkg.NewTransaction(db)
	input, _ := convertor.NewRegisterByEmailConvertor(c)

	output, _ := usecase.NewRegisterByEmailUsecase(
		ctx, tx,
		repository.NewRegisterEmailVerifyTokenRepository(db),
		repository.NewUserRepository(db),
	).Exec(input)
	presenter.RegisterByEmailPresenterExec(c, output)
}

func VerifyRegisterEmail(c *gin.Context) {
	db := connection.DBConnect()
	tx := gormpkg.NewTransaction(db)
	input := convertor.NewVerifyRegisterEmailConvertor(c)

	output, _ := usecase.NewVerifyRegisterEmailUsecase(
		tx,
		repository.NewRegisterEmailVerifyTokenRepository(db),
		repository.NewUserRepository(db),
	).Exec(input)
	presenter.VerifyRegisterEmailPresenterExec(c, output)
}

func LoginByEmail(c *gin.Context) {
	db := connection.DBConnect()
	ctx := context.Background()
	tx := gormpkg.NewTransaction(db)
	input, _ := convertor.NewLoginByEmailConvertor(c)

	output, _ := usecase.NewLoginByEmailUsecase(
		ctx, tx, repository.NewUserRepository(db),
	).Exec(input)
	presenter.LoginByEmailExec(c, output)
}
