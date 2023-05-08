package controller

import (
	"go-app/src/application/usecase"

	"github.com/gin-gonic/gin"
)

func RegisterByEmail() gin.HandlerFunc {
	return usecase.RegisterByEmail()
}

func LoginByEmail() gin.HandlerFunc {
	return usecase.LoginByEmail()
}
