package user

import (
	"syloria-demo/domain"
	userHandler "syloria-demo/rest/handler/user"
)

type Service interface {
	userHandler.Service
}

type UserRepo interface {
	Create(user domain.User) (*domain.User, error)
	Find(email, pass string) (*domain.User, error)
	//Get(userID int) (*User, error)
	//List() ([]*User, error)
	//Delete(userID int) error
	//Update(user User) (*User, error)
}
