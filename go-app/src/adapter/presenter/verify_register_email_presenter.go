package presenter

import (
	"go-app/src/application/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func VerifyRegisterEmailPresenterExec(c *gin.Context, o *usecase.VerifyRegisterEmailUsecaseOutput) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
