package users

import (
	"context"
	"errors"
	"github.com/dikopylov/highload-architect/internal/model/auth"
	"github.com/dikopylov/highload-architect/internal/model/types"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(ctx context.Context, user *User) error
	Login(ctx context.Context, id types.UserID, rawPassword string) (auth.Token, error)
	GetUserByID(ctx context.Context, id types.UserID) (*User, error)
}

type implService struct {
	repository  Repository
	authStorage auth.Storage
}

func NewService(repository Repository, authStorage auth.Storage) Service {
	return &implService{
		repository:  repository,
		authStorage: authStorage,
	}
}

func (s *implService) Register(ctx context.Context, user *User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hash)
	return s.repository.CreateUser(ctx, user)
}

func (s *implService) Login(ctx context.Context, id types.UserID, rawPassword string) (auth.Token, error) {
	user, err := s.repository.GetUserByID(ctx, id)
	if err != nil {
		return auth.EmptyToken, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(rawPassword))
	if err != nil {
		return auth.EmptyToken, err
	}

	return s.authStorage.CreateToken(user.ID), nil
}

func (s *implService) GetUserByID(ctx context.Context, id types.UserID) (*User, error) {
	if id.IsNil() {
		return nil, errors.New("id must be passed")
	}
	return s.repository.GetUserByID(ctx, id)
}
