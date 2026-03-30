package repository

import (
	"github.com/hudayberdipolat/golang-crud/internal/model"
	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{
		db: db,
	}
}

func (userRepo *userRepositoryImpl) GetAllUsers() ([]model.User, error) {
	panic("user repo")
}

func (userRepo *userRepositoryImpl) GetOneUser(userID uint) (*model.User, error) {
	panic("user repo")
}

func (userRepo *userRepositoryImpl) CreateUser(user model.User) error {
	panic("user repo")
}

func (userRepo *userRepositoryImpl) UpdateUser(userID uint, user model.User) error {
	panic("user repo")
}

func (userRepo *userRepositoryImpl) DeleteUser(userID uint) error {
	panic("user repo")
}
