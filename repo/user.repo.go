package repo

import (
	"github.com/adhaufa/nfs/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	InsertUser(user entity.User) (entity.User, error)
	UpdateUser(user entity.User) (entity.User, error)
	VerifyCredential(email string, password string) error
	IsDuplicateEmail(email string) error
	FindByEmail(email string) (entity.User, error)
	ProfileUser(userID int) (entity.User, error)
}

type userRepo struct {
	connection *gorm.DB
}

func NewUserRepo(connection *gorm.DB) UserRepository {
	return &userRepo{
		connection: connection,
	}
}

func (c *userRepo) InsertUser(user entity.User) (entity.User, error) {
	return entity.User{}, nil
}

func (c *userRepo) UpdateUser(user entity.User) (entity.User, error) {
	return entity.User{}, nil
}

func (c *userRepo) VerifyCredential(email string, password string) error {
	return nil
}

func (c *userRepo) IsDuplicateEmail(email string) error {
	return nil
}

func (c *userRepo) FindByEmail(email string) (entity.User, error) {
	return entity.User{}, nil
}

func (c *userRepo) ProfileUser(userID int) (entity.User, error) {
	return entity.User{}, nil
}
