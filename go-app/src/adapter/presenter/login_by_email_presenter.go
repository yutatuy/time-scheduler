package presenter

import (
	"go-app/src/application/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginByEmailExec(c *gin.Context, o *usecase.LoginByEmailUsecaseOutput) {
	c.JSON(http.StatusOK, gin.H{"jwt": o.Token})
}
