package convertor

import (
	"fmt"
	"go-app/src/application/usecase"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type VerifyRegisterEmailRequest struct {
	Token string `validate:"required"`
}

func NewVerifyRegisterEmailConvertor(c *gin.Context) *usecase.VerifyRegisterEmailUsecaseInput {
	var req VerifyRegisterEmailRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		panic(err)
	}

	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Printf("Error: %s\n", err)
		}
		panic(err)
	}

	return &usecase.VerifyRegisterEmailUsecaseInput{
		Token: req.Token,
	}
}
