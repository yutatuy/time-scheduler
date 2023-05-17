package repository

import (
	"crypto/rand"
	"encoding/hex"
)

func NewTokenCreator() TokenCreator {
	return &TokenCreatorImpl{}
}

type TokenCreator interface {
	Create() string
}

type TokenCreatorImpl struct {
}

func (t *TokenCreatorImpl) Create() string {
	length := 16
	randomBytes := make([]byte, length)
	rand.Read(randomBytes)
	randomString := hex.EncodeToString(randomBytes)
	return randomString
}
