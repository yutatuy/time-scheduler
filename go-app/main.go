package main

import (
	"go-app/src/adapter/controller"
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
*/

func main() {
	r := gin.Default()
	r.POST("/api/user/register", controller.RegisterByEmail())
	r.POST("/api/user/login", controller.LoginByEmail())

	if err := r.Run(":8000"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
