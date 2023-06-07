package convertor

import (
	"fmt"
	"go-app/src/application/usecase"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type verifyRegisterEmailRequest struct {
	Token string `validate:"required"`
}

type verifyRegisterEmailConvertor struct {
	c *gin.Context
}

func NewVerifyRegisterEmailConvertor(c *gin.Context) *verifyRegisterEmailConvertor {
	return &verifyRegisterEmailConvertor{
		c: c,
	}
}

func (con *verifyRegisterEmailConvertor) Exec() (*usecase.VerifyRegisterEmailUsecaseInput, error) {

	var req verifyRegisterEmailRequest
	if err := con.c.ShouldBindJSON(&req); err != nil {
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

	return &usecase.VerifyRegisterEmailUsecaseInput{
		Token: req.Token,
	}, nil
}
