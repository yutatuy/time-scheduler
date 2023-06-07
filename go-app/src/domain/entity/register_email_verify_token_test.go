package entity_test

import (
	"go-app/src/domain/entity"
	"testing"
	"time"
	_ "unsafe"
)

func TestCheckExpired(t *testing.T) {
	// given
	testCases := []struct {
		expiredAt time.Time
		answer    bool
	}{
		{
			expiredAt: time.Now(),
			answer:    true,
		},
		{
			expiredAt: time.Now().Add(time.Hour),
			answer:    false,
		},
	}

	for _, tc := range testCases {
		verifyToken := &entity.RegisterEmailVerifyToken{
			ExpiredAt: tc.expiredAt,
		}
		// when,then
		if verifyToken.CheckExpired() != tc.answer {
			t.Errorf("Error: CheckExpired")
		}
	}

}
