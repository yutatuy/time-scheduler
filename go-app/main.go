package main

import (
	"go-app/src/adapter/controller"
	"go-app/src/infrastructure/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

/**
  User
 		- name
		- email
		- password
	Schedule
 		- user_id
		- name
	Place
 		- user_id
		- schedule_id
		- name
		- position
		- time
		- before_place_id(default=0)

	- パスワード変更機能
		- {password, newPassword}
		- 過去のものがあっているか検証
	- メールアドレス変更機能
		- {newEmail}
		- 2段階認証メールを送信
		- {verifyToken} ※メールに送られたトークン
		- メアドを更新
	- パスワードを忘れた用の機能
		- {email}
		- メール送信(そのメールアドレスが存在するかチェック)
		- {verifyToken, newPassword} ※メールに送られたトークン
		- パスワード更新
*/

func main() {
	r := gin.Default()
	r.POST("/api/user/register", controller.RegisterByEmail)
	r.POST("/api/user/login", controller.LoginByEmail)
	r.POST("/api/user/verify-register-email-token", controller.VerifyRegisterEmail)

	authGroup := r.Group("")
	authGroup.Use(middleware.AuthMiddleware)
	authGroup.GET("/api/user/schedules", controller.FetchScheduleList())

	if err := r.Run(":8000"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
