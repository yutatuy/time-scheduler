package usecase

import (
	"fmt"
	"go-app/src/domain/repository"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

type LoginByEmailRequest struct {
	Email    string
	Password string
}

func LoginByEmail() gin.HandlerFunc {

	return func(c *gin.Context) {
		var req LoginByEmailRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userRepository := repository.NewUserRepository()
		user, err := userRepository.FindByEmail(req.Email)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
		}

		claims := jwt.MapClaims{
			"user_id": user.Id,
			"exp":     time.Now().Add(time.Hour * 24).Unix(),
		}
		godotenv.Load()
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
		fmt.Println("tokenString:", tokenString)

		c.JSON(http.StatusOK, gin.H{"jwt": tokenString})
	}
}
