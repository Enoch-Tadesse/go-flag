package domain

import (
	"time"
)

type FlagSegment struct {
	ID        string
	FlagID    string
	SegID     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type FlagUser struct {
	ID        string
	FlagID    string
	UserID    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type Flag struct {
	ID        string
	Name      string
	IsActive  bool
	IsRollout bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type IFlagUsecase interface {
	CreateFlag(flag Flag) error
	GetAllFlags() ([]Flag, error)
	GetFlag(name string) (*Flag, error)
	UpdateStatus(name string, isActive bool) error
	DeleteFlag(name string) error
}

type IFlagRepository interface {
	Create(flag *Flag) error
	GetAll() ([]Flag, error)
	GetByName(name string) (*Flag, error)
	UpdateStatus(name string, isActive bool) error
	Delete(name string) error
}
