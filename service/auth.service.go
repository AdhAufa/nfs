package service

import (
	"github.com/adhaufa/nfs/entity"
	"github.com/adhaufa/nfs/repo"
)

type AuthService interface {
	VerifyCredential(email string, password string) error
	// CreateUser(user dto.) entity.User
	FindByEmail(email string) (*entity.User, error)
	IsDuplicateEmail(email string) bool
}

type authService struct {
	userRepo repo.UserRepository
}

func NewAuthService(userRepo repo.UserRepository) AuthService {
	return &authService{
		userRepo: userRepo,
	}
}

func (c *authService) VerifyCredential(email string, password string) error {
	return nil
}

func (c *authService) FindByEmail(email string) (*entity.User, error) {
	return nil, nil
}

func (c *authService) IsDuplicateEmail(email string) bool {
	return false
}
