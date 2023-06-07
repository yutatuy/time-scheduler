package usecase

import (
	"context"
	application_error "go-app/src/application/error"
	"go-app/src/domain/entity"
	"go-app/src/domain/factory"
	"go-app/src/domain/repository"
	"go-app/src/domain/shared"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

type RegisterByEmailUsecase interface {
	Exec(i *RegisterByEmailUsecaseInput) (*RegisterByEmailUsecaseOutput, error)
}

type registerByEmailUsecase struct {
	ctx                                context.Context
	tx                                 shared.Transaction
	registerEmailVerifyTokenRepository repository.RegisterEmailVerifyTokenRepository
	userRepository                     repository.UserRepository
}

type RegisterByEmailUsecaseInput struct {
	Email    string
	Password string
}

type RegisterByEmailUsecaseOutput struct {
}

func NewRegisterByEmailUsecase(ctx context.Context, tx shared.Transaction, rr repository.RegisterEmailVerifyTokenRepository,
	ur repository.UserRepository) RegisterByEmailUsecase {
	return &registerByEmailUsecase{
		ctx:                                ctx,
		tx:                                 tx,
		registerEmailVerifyTokenRepository: rr,
		userRepository:                     ur,
	}
}

func (u *registerByEmailUsecase) Exec(i *RegisterByEmailUsecaseInput) (*RegisterByEmailUsecaseOutput, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(i.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, &application_error.APIError{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	txErr := u.tx.DoInTx(u.ctx, func(ctx context.Context) error {
		user, err := u.userRepository.Create(ctx, entity.User{
			Email:    i.Email,
			Password: hashedPassword,
		})

		if err != nil {
			return &application_error.APIError{Code: http.StatusInternalServerError, Message: err.Error()}
		}

		registerEmailVerifyToken, err := u.registerEmailVerifyTokenRepository.Create(ctx, factory.NewRegisterEmailVerifyToken(user.ID, user.Email))
		if err != nil {
			return &application_error.APIError{Code: http.StatusInternalServerError, Message: err.Error()}
		}

		godotenv.Load()
		email := &entity.Email{
			From:      os.Getenv("SUPPORT_ADDRESS"),
			Receivers: []string{user.Email},
			Subject:   "会員登録",
			Body:      registerEmailVerifyToken.Token,
		}
		emailSender := repository.NewEmailSender()
		emailSender.Send(email)
		return nil
	})

	if txErr != nil {
		return nil, &application_error.APIError{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	return &RegisterByEmailUsecaseOutput{}, nil
}
