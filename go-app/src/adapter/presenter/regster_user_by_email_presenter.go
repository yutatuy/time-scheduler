package presenter

import (
	"go-app/src/application/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterByEmailPresenterExec(c *gin.Context, o *usecase.RegisterByEmailUsecaseOutput) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
