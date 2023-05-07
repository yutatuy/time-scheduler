package usecase

import (
	"fmt"
	entity "go-app/src/domain/entity/user"
	"go-app/src/domain/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RegisterUserByEmailRequest struct {
	Email    string
	Password string
}

func RegisterUserByEmail() gin.HandlerFunc {

	return func(c *gin.Context) {
		var req RegisterUserByEmailRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		fmt.Println(req.Password)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not"})
		}

		userRepository := repository.NewUserRepository()
		user, err := userRepository.CreateByEmail(entity.User{
			Email:    req.Email,
			Password: hashedPassword,
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(user)

		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
}
