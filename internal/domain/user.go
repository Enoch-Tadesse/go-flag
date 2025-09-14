package domain

import (
	"time"
)

type User struct {
	ID        string
	Email     string
	Country   string
	Paid      bool
	SignUp    time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type IUserUsecase interface {
	RegisterUser(user *User) error
	GetUserByID(id string) (*User, error)
	UpdateUser(user *User) error
	DeleteUser(id string) error
}


type IUserRepository interface {
    Create(user *User) error
    GetByID(id string) (*User, error)
    Update(user *User) error
    Delete(id string) error
}
