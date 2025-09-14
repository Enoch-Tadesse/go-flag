// internal/usecase/user_usecase.go
package usecase

import "github.com/Enoch-Tadesse/go-flag/internal/domain"

type UserUsecase struct {
	repo domain.IUserRepository
}

// DeleteUser implements domain.IUserUsecase.
func (u *UserUsecase) DeleteUser(id string) error {
	panic("unimplemented")
}

// GetUserByID implements domain.IUserUsecase.
func (u *UserUsecase) GetUserByID(id string) (*domain.User, error) {
	panic("unimplemented")
}

// RegisterUser implements domain.IUserUsecase.
func (u *UserUsecase) RegisterUser(user *domain.User) error {
	panic("unimplemented")
}

// UpdateUser implements domain.IUserUsecase.
func (u *UserUsecase) UpdateUser(user *domain.User) error {
	panic("unimplemented")
}

func NewUserUsecase(r domain.IUserRepository) domain.IUserUsecase {
	return &UserUsecase{repo: r}
}
