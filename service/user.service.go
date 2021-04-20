package service

import (
	"errors"
	"log"

	"github.com/adhaufa/nfs/dto"
	"github.com/adhaufa/nfs/repo"
	_user "github.com/adhaufa/nfs/service/user"
	"github.com/mashingan/smapping"
	"gorm.io/gorm"
)

type UserService interface {
	CreateUser(registerRequest dto.RegisterRequest) (*_user.UserResponse, error)
	FindUserByEmail(email string) (*_user.UserResponse, error)
}

type userService struct {
	userRepo repo.UserRepository
}

func NewUserService(userRepo repo.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (c *userService) CreateUser(registerRequest dto.RegisterRequest) (*_user.UserResponse, error) {
	user, err := c.userRepo.FindByEmail(registerRequest.Email)

	if err == nil {
		return nil, errors.New("user already exists")
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	err = smapping.FillStruct(&user, smapping.MapFields(&registerRequest))

	if err != nil {
		log.Fatalf("Failed map %v", err)
		return nil, err
	}

	user, _ = c.userRepo.InsertUser(user)
	res := _user.NewUserResponse(user)
	return &res, nil

}

func (c *userService) FindUserByEmail(email string) (*_user.UserResponse, error) {
	user, err := c.userRepo.FindByEmail(email)

	if err != nil {
		return nil, err
	}

	userResponse := _user.NewUserResponse(user)
	return &userResponse, nil
}
