package mapper

import "github.com/Enoch-Tadesse/go-flag/internal/domain"
import "github.com/Enoch-Tadesse/go-flag/internal/delivery/dto"

// --- DTO -> Domain

func ToDomainSegment(req dto.CreateSegmentRequest) domain.Segment {
	rules := make([]domain.SegmentRule, len(req.Rules))
	for i, r := range req.Rules {
		rules[i] = domain.SegmentRule{
			Attribute: r.Attribute,
			Operator:  r.Operator,
			Value:     r.Value,
		}
	}

	return domain.Segment{
		Name:  req.Name,
		Rules: rules,
	}
}

// --- Domain -> DTO

func FromDomainSegment(s domain.Segment) dto.SegmentResponse {
	rules := make([]dto.SegmentRuleDTO, len(s.Rules))
	for i, r := range s.Rules {
		rules[i] = dto.SegmentRuleDTO{
			Attribute: r.Attribute,
			Operator:  r.Operator,
			Value:     r.Value,
		}
	}

	return dto.SegmentResponse{
		ID:        s.ID,
		Name:      s.Name,
		Rules:     rules,
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
	}
}

func FromDomainSegments(segments []domain.Segment) []dto.SegmentListResponse {
	result := make([]dto.SegmentListResponse, len(segments))
	for i, s := range segments {
		result[i] = dto.SegmentListResponse{
			ID:   s.ID,
			Name: s.Name,
		}
	}
	return result
}
