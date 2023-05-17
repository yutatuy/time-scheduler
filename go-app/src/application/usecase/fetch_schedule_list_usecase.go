package usecase

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type FetchScheduleListRequest struct{}

func FetchScheduleList() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
}
