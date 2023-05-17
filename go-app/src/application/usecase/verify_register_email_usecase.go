package usecase

import (
	"errors"
	"go-app/src/domain/repository"
	"go-app/src/domain/shared"
)

type VerifyRegisterEmailUsecase interface {
	Exec(i *VerifyRegisterEmailUsecaseInput) (*VerifyRegisterEmailUsecaseOutput, error)
}

type verifyRegisterEmailUsecase struct {
	tx                                 shared.Transaction
	registerEmailVerifyTokenRepository repository.RegisterEmailVerifyTokenRepository
	userRepository                     repository.UserRepository
}

type VerifyRegisterEmailUsecaseInput struct {
	Token string
}

type VerifyRegisterEmailUsecaseOutput struct {
}

func NewVerifyRegisterEmailUsecase(tx shared.Transaction, rr repository.RegisterEmailVerifyTokenRepository, ur repository.UserRepository,
) VerifyRegisterEmailUsecase {
	return &verifyRegisterEmailUsecase{
		tx:                                 tx,
		registerEmailVerifyTokenRepository: rr,
		userRepository:                     ur,
	}
}

func (u *verifyRegisterEmailUsecase) Exec(i *VerifyRegisterEmailUsecaseInput) (*VerifyRegisterEmailUsecaseOutput, error) {
	token, err := u.registerEmailVerifyTokenRepository.FindByToken(i.Token)
	if err != nil {
		return nil, err
	}

	if token.CheckExpired() {
		return nil, errors.New("Token Expired!")
	}

	user, err := u.userRepository.FindByID(token.UserID)
	if err != nil {
		return nil, err
	}

	user.VerifyEmail()
	_, updateErr := u.userRepository.Update(user)
	if updateErr != nil {
		return nil, updateErr
	}

	return &VerifyRegisterEmailUsecaseOutput{}, nil
}
