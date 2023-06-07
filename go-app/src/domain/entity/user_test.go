package entity_test

import (
	"go-app/src/domain/entity"
	"testing"
	_ "unsafe"
)

func TestSetEmailVerifiedAt(t *testing.T) {

	user := &entity.User{}
	user.SetEmailVerifiedAt()

	if user.EmailVerifiedAt == nil {
		t.Errorf("EmailVerifiedAt should not be nil")
	}
}
