package presenter

import (
	adapter_share "go-app/src/adapter/share"

	"github.com/gin-gonic/gin"
)

type verifyRegisterEmailPresenter struct {
	adapter_share.BaseAPIPresenter
}

func NewVerifyRegisterEmailPresenter(ctx *gin.Context) *verifyRegisterEmailPresenter {
	return &verifyRegisterEmailPresenter{
		BaseAPIPresenter: adapter_share.BaseAPIPresenter{
			Ctx: ctx,
		},
	}
}
