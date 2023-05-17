package convertor

import (
	"fmt"
	"go-app/src/application/usecase"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type RegisterByEmailRequest struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8,max=20"`
}

func NewRegisterByEmailConvertor(c *gin.Context) (*usecase.RegisterByEmailUsecaseInput, error) {
	var req RegisterByEmailRequest

	if err := c.ShouldBindJSON(&req); err != nil {
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

	return &usecase.RegisterByEmailUsecaseInput{
		Email:    req.Email,
		Password: req.Password,
	}, nil
}
