package request

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GetParagraphsV1RequestSuite struct {
	suite.Suite
}

func TestGetParagraphsV1RequestSuite(t *testing.T) {
	suite.Run(t, new(GetParagraphsV1RequestSuite))
}

func (s *GetParagraphsV1RequestSuite) TestNewGetParagraphsV1Request() {
	ctx := context.Background()
	httpReq := httptest.NewRequest(http.MethodGet, "/paragraphs", nil)

	req, err := NewGetParagraphsV1Request(ctx, httpReq)

	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), req)
}

func (s *GetParagraphsV1RequestSuite) TestValidate() {
	req := &GetParagraphsV1Request{}

	err := req.Validate()

	assert.NoError(s.T(), err)
}
