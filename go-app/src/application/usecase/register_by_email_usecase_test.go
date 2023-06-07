package usecase_test

import (
	"context"
	"go-app/src/application/usecase"
	"go-app/src/domain/repository"
	gormpkg "go-app/src/infrastructure/gorm"
	"go-app/src/infrastructure/gorm/model"
	"go-app/src/infrastructure/testutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterByEmailUsecase(t *testing.T) {
	// given
	ctx := context.Background()
	db := testutil.SetUp()
	tx := gormpkg.NewTransaction(db)
	testutil.RefreshTables(db, []interface{}{&model.GormUser{}, &model.GormRegisterEmailVerifyToken{}})

	input := &usecase.RegisterByEmailUsecaseInput{
		Email:    "test1@info.com",
		Password: "password",
	}
	revtr := repository.NewRegisterEmailVerifyTokenRepository(db)
	ur := repository.NewUserRepository(db)
	usecase := usecase.NewRegisterByEmailUsecase(
		ctx, tx, revtr, ur,
	)

	// when
	usecase.Exec(input)

	// then
	// Userが保存されていること
	user, _ := ur.FindByEmail(input.Email)
	assert.NotNil(t, user, "User should not be nil")

	// RegisterEmailVerifyTokenが保存されていること
	registerEmailVerifyToken, _ := revtr.FindByUserIDAndEmail(user.ID, user.Email)
	assert.Equal(t, user.ID, registerEmailVerifyToken.UserID, "UserId should be equal")
	assert.Equal(t, user.Email, registerEmailVerifyToken.Email, "Email should be equal")

}
