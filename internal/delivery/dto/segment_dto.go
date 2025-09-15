package dto

import "time"

// Request when creating a new segment
type CreateSegmentRequest struct {
	Name  string           `json:"name" binding:"required"`
	Rules []SegmentRuleDTO `json:"rules"`
}

// Response when returning full segment details
type SegmentResponse struct {
	ID        string           `json:"id"`
	Name      string           `json:"name"`
	Rules     []SegmentRuleDTO `json:"rules"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
}

// Response for lightweight segment list
type SegmentListResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Rule DTO for requests and responses
type SegmentRuleDTO struct {
	Attribute string `json:"attribute"`
	Operator  string `json:"operator"`
	Value     string `json:"value"`
}
