package user

import (
	"github.com/alwindoss/manna/internal/storage"
)

type User struct {
	UserName string
	Password string
}

type Service interface {
	IsAuthenticated(*User) bool
	Create(*User) error
}

func NewService(r storage.UserRepository) Service {
	return &userService{
		repo: r,
	}
}

type userService struct {
	repo storage.UserRepository
}

// Create implements Service
func (r *userService) Create(u *User) error {
	user := &storage.User{
		UserName: u.UserName,
		Password: u.Password,
	}
	return r.repo.Create(user)
}

// IsAuthenticated implements Service
func (r *userService) IsAuthenticated(u *User) bool {
	user := &storage.User{
		UserName: u.UserName,
		Password: u.Password,
	}
	return r.repo.IsAuthenticated(user)
}
