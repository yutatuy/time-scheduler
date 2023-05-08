package controller

import (
	"go-app/src/application/usecase"

	"github.com/gin-gonic/gin"
)

func FetchScheduleList() gin.HandlerFunc {
	return usecase.FetchScheduleList()
}
