// internal/usecase/flag_usecase.go
package usecase

import (
	"github.com/Enoch-Tadesse/go-flag/internal/domain"
)

type FlagUsecase struct {
	repo domain.IFlagRepository
}

// CreateFlag implements domain.IFlagUsecase.
func (f *FlagUsecase) CreateFlag(flag domain.Flag) error {
	panic("unimplemented")
}

// DeleteFlag implements domain.IFlagUsecase.
func (f *FlagUsecase) DeleteFlag(name string) error {
	panic("unimplemented")
}

// GetAllFlags implements domain.IFlagUsecase.
func (f *FlagUsecase) GetAllFlags() ([]domain.Flag, error) {
	panic("unimplemented")
}

// GetFlag implements domain.IFlagUsecase.
func (f *FlagUsecase) GetFlag(name string) (*domain.Flag, error) {
	panic("unimplemented")
}

// UpdateStatus implements domain.IFlagUsecase.
func (f *FlagUsecase) UpdateStatus(name string, isActive bool) error {
	panic("unimplemented")
}

func NewFlagUsecase(r domain.IFlagRepository) domain.IFlagUsecase {
	return &FlagUsecase{repo: r}
}
