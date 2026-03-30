package repository

import "github.com/hudayberdipolat/golang-crud/internal/model"

type UserRepository interface {
	GetAllUsers() ([]model.User, error)
	GetOneUser(userID uint) (*model.User, error)
	CreateUser(user model.User) error
	UpdateUser(userID uint, user model.User) error
	DeleteUser(userID uint) error
}
