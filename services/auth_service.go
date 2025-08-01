package services

import (
	"errors"

	"github.com/rijwanansari/vivaLearning/domain"
	repository "github.com/rijwanansari/vivaLearning/repositories"
	"github.com/rijwanansari/vivaLearning/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(email, password string) error
	Login(email, password string) (string, error)
}

type AuthServiceImp struct {
	UserRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) *AuthServiceImp {
	return &AuthServiceImp{UserRepo: userRepo}
}

func (s *AuthServiceImp) Register(email, password string) (*domain.User, error) {
	if existing, _ := s.UserRepo.GetByEmail(email); existing != nil {
		return nil, errors.New("email already registered")
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := &domain.User{
		Email:    email,
		Password: string(hashed),
		Role:     "user",
	}
	if err := s.UserRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}
func (s *AuthServiceImp) Login(email, password string) (string, error) {
	user, err := s.UserRepo.GetByEmail(email)
	if err != nil {
		return "", err
	}

	if user == nil {
		return "", errors.New("user not found")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid password")
	}
	return utils.GenerateJWT(user.ID)
}
