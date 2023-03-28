package storage

import "sync"

type User struct {
	UserName string
	Password string
}

type UserRepository interface {
	IsAuthenticated(*User) bool
	Create(*User) error
}

func NewInMemRepository() UserRepository {
	var repo *userRepository
	inMemOnce.Do(func() {
		repo = &userRepository{
			userDB:    map[string]*User{},
			sessionDB: map[string]string{},
		}
	})
	return repo
}

var inMemOnce sync.Once

type userRepository struct {
	// user DB stores user name as key and user object as value
	userDB map[string]*User
	// session DB stores user name as key and session id as value
	sessionDB map[string]string
}

// Create implements UserRepository
func (r *userRepository) Create(u *User) error {
	r.userDB[u.UserName] = u
	return nil
}

// IsAuthenticated implements UserRepository
func (r *userRepository) IsAuthenticated(u *User) bool {
	isAuthenticated := false
	user, found := r.userDB[u.UserName]
	if found {
		if user.Password == u.Password {
			isAuthenticated = true
		}
	}
	return isAuthenticated
}
