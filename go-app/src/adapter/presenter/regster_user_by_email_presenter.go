package presenter

import (
	adapter_share "go-app/src/adapter/share"

	"github.com/gin-gonic/gin"
)

type registerByEmailPresenter struct {
	adapter_share.BaseAPIPresenter
}

func NewRegisterByEmailPresenter(ctx *gin.Context) *registerByEmailPresenter {
	return &registerByEmailPresenter{
		BaseAPIPresenter: adapter_share.BaseAPIPresenter{
			Ctx: ctx,
		},
	}
}
