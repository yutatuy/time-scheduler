package convertor

import (
	"fmt"
	"go-app/src/application/usecase"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type loginByEmailRequest struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8,max=20"`
}

type loginByEmailConverter struct {
	ctx *gin.Context
}

func NewLoginByEmailConvertor(ctx *gin.Context) *loginByEmailConverter {
	return &loginByEmailConverter{
		ctx: ctx,
	}
}

func (con *loginByEmailConverter) Exec() (*usecase.LoginByEmailUsecaseInput, error) {
	var req loginByEmailRequest

	if err := con.ctx.ShouldBindJSON(&req); err != nil {
		return nil, err
	}

	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Printf("Error: %s\n", err)
		}
		return nil, err
	}

	return &usecase.LoginByEmailUsecaseInput{
		Email:    req.Email,
		Password: req.Password,
	}, nil
}
