package auth

import (
	"errors"
	"github.com/dikopylov/highload-architect/internal/model/types"
	"github.com/google/uuid"
)

type Token string

func (t Token) String() string {
	return string(t)
}

var EmptyToken Token

type UserByToken map[Token]types.UserID

var ErrUserNotFound = errors.New("user not found")

type Storage interface {
	CreateToken(inputUserID types.UserID) Token
	GetUserID(token Token) (types.UserID, error)
}

type inMemoryStorage struct {
	storage UserByToken
}

func NewInMemoryStorage() Storage {
	return &inMemoryStorage{storage: make(UserByToken)}
}

func (s *inMemoryStorage) CreateToken(inputUserID types.UserID) Token {
	for token, userID := range s.storage {
		if userID == inputUserID {
			return token
		}
	}

	token := Token(uuid.New().String())
	s.storage[token] = inputUserID

	return token
}

func (s *inMemoryStorage) GetUserID(token Token) (types.UserID, error) {
	if userID, ok := s.storage[token]; ok {
		return userID, nil
	}

	return types.NilUserID, ErrUserNotFound
}
