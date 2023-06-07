package adapter_share

import (
	application_error "go-app/src/application/error"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseAPIPresenter struct {
	Ctx *gin.Context
}

func (p *BaseAPIPresenter) Exec(o interface{}) {
	p.Ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (p *BaseAPIPresenter) Error(err error) {
	if err, ok := err.(*application_error.APIError); ok {
		p.Ctx.JSON(err.Code, gin.H{"code": err.Code, "message": err.Message})
		return
	}
	p.Ctx.JSON(http.StatusInternalServerError, err)
}
