package services

import (
	"errors"
	"time"

	"gtihub.com/rijwanansari/vivaLearning/domain"
	repository "gtihub.com/rijwanansari/vivaLearning/repositories"
	util "gtihub.com/rijwanansari/vivaLearning/utils"
)

type UserService interface {
	RegisterUser(email, password string) (*domain.User, error)
}

type userServiceImpl struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *userServiceImpl {
	return &userServiceImpl{repo: repo}
}

func (s *userServiceImpl) RegisterUser(email, password string) (*domain.User, error) {
	if existing, _ := s.repo.GetByEmail(email); existing != nil {
		return nil, errors.New("email already registered")
	}

	hashed, err := util.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Email:     email,
		Password:  hashed,
		Role:      "user",
		CreatedAt: time.Now(),
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}
