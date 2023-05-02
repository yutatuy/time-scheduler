package usecase

import (
	"fmt"
	entity "go-app/src/domain/entity/user"
	"go-app/src/domain/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUserByEmail() gin.HandlerFunc {

	return func(c *gin.Context) {
		var userEntity entity.User

		if err := c.ShouldBindJSON(&userEntity); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userEntity.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not"})
		}

		userEntity.Password = string(hashedPassword)
		userRepository := repository.NewUserRepository()
		user, err := userRepository.CreateByEmail(userEntity)
		if err != nil {
			panic(err)
		}
		fmt.Println(user)

		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
}
