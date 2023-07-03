package entity

import (
	"go-app/src/infrastructure/configs"
)

type Email struct {
	From      string
	Subject   string
	Receivers []string
	TemplateFiles []string
	TemplateVars interface{}
}

// TODO: 別ファイルに分けた方がFactoryが増えた時にみやすい
func RegisterEmailFactory (receiver string, name string) *Email {
	return &Email{
		From:      configs.Config.Common.SupportAddress,
		Receivers: []string{receiver},
		Subject:   "会員登録",
		TemplateFiles: []string{"/app/go/src/infrastructure/mail/templates/register.tpl"},
		TemplateVars: struct {
				Name string
				Mail string
			}{
				Name: name,
				Mail: receiver,
			},
	}
}
