package response

import (
	"net/http"
	"testing"
	"time"

	"github.com/gagansingh3785/typio-service/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GetParagraphsV1ResponseSuite struct {
	suite.Suite
}

func TestGetParagraphsV1ResponseSuite(t *testing.T) {
	suite.Run(t, new(GetParagraphsV1ResponseSuite))
}

func (s *GetParagraphsV1ResponseSuite) TestNewGetParagraphsV1Response() {
	now := time.Now()
	paragraph := &domain.Paragraph{
		ID:        1,
		UUID:      "test-uuid",
		Content:   "The quick brown fox jumps over the lazy dog.",
		CreatedAt: now,
		UpdatedAt: now,
	}

	resp := NewGetParagraphsV1Response(paragraph)

	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), paragraph.Content, resp.Content)
}

func (s *GetParagraphsV1ResponseSuite) TestGetStatus() {
	resp := &GetParagraphsV1Response{
		Content: "Test content",
	}

	status := resp.GetStatus()

	assert.Equal(s.T(), http.StatusOK, status)
}
