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
	GetSegment(id string) (*Segment, error)
	AddRule(name string, rule *SegmentRule) error
	DeleteSegment(id string) error
	DeleteRule(id string) error
}

type ISegmentRepository interface {
	Create(segment *Segment) error
	GetAll() ([]Segment, error)
	GetByID(id string) (*Segment, error)
	AppendRule(segmentName string, rule *SegmentRule) error
	DeleteSegment(id string) error
	RemoveRule(segmentName string, rule *SegmentRule) error
}
