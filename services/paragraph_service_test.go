package services

import (
	"context"
	"errors"
	"testing"

	"github.com/gagansingh3785/typio-service/builders"
	"github.com/gagansingh3785/typio-service/builders/mocks"
	"github.com/gagansingh3785/typio-service/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type ParagraphServiceSuite struct {
	suite.Suite

	ctx           context.Context
	mockBuilder   *mocks.ParagraphsBuilder
	service       ParagraphService
	testParagraph *domain.Paragraph
}

func (s *ParagraphServiceSuite) SetupTest() {
	s.ctx = context.Background()
	s.mockBuilder = &mocks.ParagraphsBuilder{}
	s.service = NewParagraphService(s.mockBuilder)
	s.testParagraph = &domain.Paragraph{
		Content: "This is a sample paragraph!",
	}
}

func (s *ParagraphServiceSuite) TestGetRandomParagraph_Success() {
	respChan := make(chan builders.ParagraphResponseAsync, 1)
	respChan <- builders.ParagraphResponseAsync{
		Paragraph: s.testParagraph,
		Error:     nil,
	}
	close(respChan)

	s.mockBuilder.On("GetAsycChannel", s.ctx).Return(respChan).Once()
	s.mockBuilder.On("Build", s.ctx, mock.AnythingOfType("builders.ParagraphResponseAsync")).
		Return(s.testParagraph, nil).Once()

	paragraph, err := s.service.GetRandomParagraph(s.ctx)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), s.testParagraph, paragraph)
	s.mockBuilder.AssertExpectations(s.T())
}

func (s *ParagraphServiceSuite) TestGetRandomParagraph_Error() {
	expectedErr := errors.New("builder failed")
	respChan := make(chan builders.ParagraphResponseAsync, 1)
	respChan <- builders.ParagraphResponseAsync{
		Paragraph: nil,
		Error:     expectedErr,
	}
	close(respChan)

	s.mockBuilder.On("GetAsycChannel", s.ctx).Return(respChan).Once()
	s.mockBuilder.On("Build", s.ctx, mock.AnythingOfType("builders.ParagraphResponseAsync")).
		Return((*domain.Paragraph)(nil), expectedErr).Once()

	paragraph, err := s.service.GetRandomParagraph(s.ctx)
	assert.Error(s.T(), err)
	assert.Nil(s.T(), paragraph)
	s.mockBuilder.AssertExpectations(s.T())
}

func TestParagraphServiceSuite(t *testing.T) {
	suite.Run(t, new(ParagraphServiceSuite))
}
