package domain

import (
	"time"
)

type Segment struct {
	ID        string
	Name      string
	Rules     []SegmentRule
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type SegmentRule struct {
	ID        string
	SegID     string
	Attribute string
	Operator  string
	Value     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type ISegmentUsecase interface {
	CreateSegment(segment *Segment) error
	GetAllSegments() ([]Segment, error)
	GetSegment(name string) (*Segment, error)
	AddRule(name string, rule *SegmentRule) error
	DeleteSegment(name string) error
}

type ISegmentRepository interface {
	Create(segment *Segment) error
	GetAll() ([]Segment, error)
	GetByName(name string) (*Segment, error)
	AddRule(segmentName string, rule *SegmentRule) error
	RemoveRule(segmentName string, rule *SegmentRule) error
	Delete(name string) error
}
