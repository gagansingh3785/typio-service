package handlers

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gagansingh3785/typio-service/domain"
	"github.com/gagansingh3785/typio-service/registry"
	svcmocks "github.com/gagansingh3785/typio-service/services/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type GetParagraphsV1HandlerSuite struct {
	suite.Suite
	mockService *svcmocks.ParagraphService

	handler http.HandlerFunc
}

func (s *GetParagraphsV1HandlerSuite) SetupTest() {
	s.mockService = &svcmocks.ParagraphService{}
	svcRegistry := &registry.ServiceRegistry{
		ParagraphService: s.mockService,
	}
	s.handler = ParagraphsV1Handler(svcRegistry)
}

func (s *GetParagraphsV1HandlerSuite) TestSuccess() {
	paragraph := &domain.Paragraph{
		Content: "test content",
	}
	s.mockService.On("GetRandomParagraph", mock.Anything).Return(paragraph, nil)

	req, err := http.NewRequest("GET", "/v1/paragraphs", nil)
	assert.NoError(s.T(), err)

	recorder := httptest.NewRecorder()

	s.handler(recorder, req)

	resp := recorder.Result()
	respBody, err := io.ReadAll(resp.Body)
	assert.NoError(s.T(), err)
	respBodyString := string(respBody)

	assert.Equal(s.T(), http.StatusOK, resp.StatusCode)
	assert.JSONEq(s.T(), `{"content": "test content"}`, respBodyString)

	s.mockService.AssertExpectations(s.T())
}

func (s *GetParagraphsV1HandlerSuite) TestBuilderError() {
	s.mockService.
		On("GetRandomParagraph", mock.Anything).Return(nil, errors.New("error"))

	req, err := http.NewRequest("GET", "/v1/paragraphs", nil)
	assert.NoError(s.T(), err)
	recorder := httptest.NewRecorder()

	s.handler(recorder, req)

	resp := recorder.Result()
	assert.Equal(s.T(), http.StatusInternalServerError, resp.StatusCode)
	s.mockService.AssertExpectations(s.T())
}

func TestGetParagraphsV1HandlerSuite(t *testing.T) {
	suite.Run(t, new(GetParagraphsV1HandlerSuite))
}
