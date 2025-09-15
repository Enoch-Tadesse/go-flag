// internal/usecase/segment_usecase.go
package usecase

import (
	"github.com/Enoch-Tadesse/go-flag/internal/domain"
)

type SegmentUsecase struct {
	repo domain.ISegmentRepository
}

// DeleteRule implements domain.ISegmentUsecase.
func (s *SegmentUsecase) DeleteRule(id string) error {
	panic("unimplemented")
}

// AddRule implements domain.ISegmentUsecase.
func (s *SegmentUsecase) AddRule(name string, rule *domain.SegmentRule) error {
	panic("unimplemented")
}

// CreateSegment implements domain.ISegmentUsecase.
func (s *SegmentUsecase) CreateSegment(segment *domain.Segment) error {
	panic("unimplemented")
}

// DeleteSegment implements domain.ISegmentUsecase.
func (s *SegmentUsecase) DeleteSegment(name string) error {
	panic("unimplemented")
}

// GetAllSegments implements domain.ISegmentUsecase.
func (s *SegmentUsecase) GetAllSegments() ([]domain.Segment, error) {
	panic("unimplemented")
}

// GetSegment implements domain.ISegmentUsecase.
func (s *SegmentUsecase) GetSegment(id string) (*domain.Segment, error) {
	panic("unimplemented")
}

func NewSegmentUsecase(r domain.ISegmentRepository) domain.ISegmentUsecase {
	return &SegmentUsecase{repo: r}
}
