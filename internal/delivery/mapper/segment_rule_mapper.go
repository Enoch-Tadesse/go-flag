package mapper

import (
	"github.com/Enoch-Tadesse/go-flag/internal/delivery/dto"
	"github.com/Enoch-Tadesse/go-flag/internal/domain"
)

// DTO -> Domain
func ToDomainSegmentRule(req dto.CreateSegmentRuleRequest, segmentID string) domain.SegmentRule {
	return domain.SegmentRule{
		SegID:     segmentID,
		Attribute: req.Attribute,
		Operator:  req.Operator,
		Value:     req.Value,
	}
}

// Domain -> DTO
func FromDomainSegmentRule(r domain.SegmentRule) dto.SegmentRuleResponse {
	return dto.SegmentRuleResponse{
		ID:        r.ID,
		SegID:     r.SegID,
		Attribute: r.Attribute,
		Operator:  r.Operator,
		Value:     r.Value,
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
	}
}
