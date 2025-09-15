package controllers

import (
	"net/http"
	"strings"
	"time"

	"github.com/Enoch-Tadesse/go-flag/internal/delivery/dto"
	"github.com/Enoch-Tadesse/go-flag/internal/delivery/mapper"
	"github.com/Enoch-Tadesse/go-flag/internal/usecase"
	"github.com/gin-gonic/gin"
)

type SegmentController struct {
	Timeout        time.Duration
	SegmentUsecase usecase.SegmentUsecase
}

func NewSegmentController(timeout time.Duration, segmentUsecase usecase.SegmentUsecase) *SegmentController {
	return &SegmentController{Timeout: timeout, SegmentUsecase: segmentUsecase}
}

// CreateSegment handles the creation of a new segment with rules.
func (sc *SegmentController) CreateSegment(c *gin.Context) {

	// Parse and validate the request body
	var body dto.CreateSegmentRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	domainBody := mapper.ToDomainSegment(body)

	err := sc.SegmentUsecase.CreateSegment(&domainBody)
	if err != nil {
		switch {
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"err": "somehting went wrong"})
		}
		return
	}

	c.JSON(http.StatusCreated, domainBody)
}

func (sc *SegmentController) GetAllSegments(c *gin.Context) {
	segments, err := sc.SegmentUsecase.GetAllSegments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, segments)
}

func (sc *SegmentController) GetSegmentByName(c *gin.Context) {
	id := c.Param("id")
	segment, err := sc.SegmentUsecase.GetSegment(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, segment)
}

func (sc *SegmentController) DeleteSegment(c *gin.Context) {
	id := c.Param("id")
	err := sc.SegmentUsecase.DeleteSegment(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "")
}

// DeleteRule deletes a segment rule by its UUID.
// The rule ID is extracted from the URL path parameter `id`.
func (sc *SegmentController) DeleteRule(c *gin.Context) {
	id := c.Param("id")
	err := sc.SegmentUsecase.DeleteRule(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "")
}

// AppendRule handles adding a new rule to an existing segment.
// It expects JSON input with segment Name, attribute, operator, and value.
// On success, it returns the newly created rule.
func (sc *SegmentController) AppendRule(c *gin.Context) {

	seg_id := c.Param("id")

	// Input structure expects segment name to find the segment
	var body dto.CreateSegmentRuleRequest

	// Decode JSON request
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body", "detail": err.Error()})
	}

	// trim the datas
	body.Attribute = strings.TrimSpace(body.Attribute)
	body.Operator = strings.TrimSpace(body.Operator)
	body.Value = strings.TrimSpace(body.Value)

	rule := mapper.ToDomainSegmentRule(body, seg_id)

	// Find segment by id, scan fields explicitly
	err := sc.SegmentUsecase.AddRule(seg_id, &rule)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to append rule", "detail": err.Error()})
		return
	}
	c.JSON(http.StatusOK, rule)

}
