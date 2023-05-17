package usecase

import (
	"context"
	"fmt"
	"go-app/src/domain/repository"
	"go-app/src/domain/shared"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

type LoginByEmailRequest struct {
	Email    string
	Password string
}

type LoginByEmailUsecaseInput struct {
	Email    string
	Password string
}

type LoginByEmailUsecaseOutput struct {
	Token string
}

type LoginByEmailUsecase interface {
	Exec(i *LoginByEmailUsecaseInput) (*LoginByEmailUsecaseOutput, error)
}

type loginByEmailUsecase struct {
	userRepository repository.UserRepository
	ctx            context.Context
	tx             shared.Transaction
}

func NewLoginByEmailUsecase(ctx context.Context, tx shared.Transaction, ur repository.UserRepository) LoginByEmailUsecase {
	return &loginByEmailUsecase{
		userRepository: ur,
		ctx:            ctx,
		tx:             tx,
	}
}

func (u *loginByEmailUsecase) Exec(i *LoginByEmailUsecaseInput) (*LoginByEmailUsecaseOutput, error) {

	user, err := u.userRepository.FindByEmail(i.Email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(i.Password)); err != nil {
		return nil, err
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	godotenv.Load()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	fmt.Println("tokenString:", tokenString)

	return &LoginByEmailUsecaseOutput{Token: tokenString}, nil

}
