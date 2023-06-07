package presenter

import (
	adapter_share "go-app/src/adapter/share"

	"github.com/gin-gonic/gin"
)

type loginByEmailPresenter struct {
	adapter_share.BaseAPIPresenter
}

func NewLoginByEmailPresenter(ctx *gin.Context) *loginByEmailPresenter {
	return &loginByEmailPresenter{
		BaseAPIPresenter: adapter_share.BaseAPIPresenter{
			Ctx: ctx,
		},
	}
}
