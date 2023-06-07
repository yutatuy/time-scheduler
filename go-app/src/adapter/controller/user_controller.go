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
	presenter := presenter.NewRegisterByEmailPresenter(c)
	convertor := convertor.NewRegisterByEmailConvertor(c)
	usecase := usecase.NewRegisterByEmailUsecase(
		ctx, tx,
		repository.NewRegisterEmailVerifyTokenRepository(db),
		repository.NewUserRepository(db),
	)

	input, err := convertor.Exec()
	if err != nil {
		presenter.Error(err)
		return
	}
	output, err := usecase.Exec(input)
	if err != nil {
		presenter.Error(err)
		return
	}
	presenter.Exec(output)
}

func VerifyRegisterEmail(c *gin.Context) {
	db := connection.DBConnect()
	tx := gormpkg.NewTransaction(db)
	presenter := presenter.NewVerifyRegisterEmailPresenter(c)
	convertor := convertor.NewVerifyRegisterEmailConvertor(c)
	usecase := usecase.NewVerifyRegisterEmailUsecase(
		tx,
		repository.NewRegisterEmailVerifyTokenRepository(db),
		repository.NewUserRepository(db),
	)

	input, err := convertor.Exec()
	if err != nil {
		presenter.Error(err)
		return
	}
	output, err := usecase.Exec(input)
	if err != nil {
		presenter.Error(err)
		return
	}
	presenter.Exec(output)
}

func LoginByEmail(c *gin.Context) {
	db := connection.DBConnect()
	ctx := context.Background()
	tx := gormpkg.NewTransaction(db)
	convertor := convertor.NewLoginByEmailConvertor(c)
	presenter := presenter.NewLoginByEmailPresenter(c)
	usecase := usecase.NewLoginByEmailUsecase(
		ctx, tx, repository.NewUserRepository(db),
	)

	input, err := convertor.Exec()
	if err != nil {
		presenter.Error(err)
		return
	}
	output, err := usecase.Exec(input)
	if err != nil {
		presenter.Error(err)
		return
	}
	presenter.Exec(output)
}
