package builders

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/gagansingh3785/typio-service/domain"
	svcerrors "github.com/gagansingh3785/typio-service/errors"
	"github.com/gagansingh3785/typio-service/repository/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ParagraphsBuilderSuite struct {
	suite.Suite

	ctx context.Context

	mockRepo      *mocks.Repository
	builder       ParagraphsBuilder
	testParagraph *domain.Paragraph
}

func (s *ParagraphsBuilderSuite) SetupTest() {
	s.ctx = context.Background()
	s.mockRepo = &mocks.Repository{}
	s.builder = NewParagraphsBuilder(s.mockRepo)
	s.testParagraph = &domain.Paragraph{
		Content: "Hello world!",
	}
}

func (s *ParagraphsBuilderSuite) TestGetAsyncChannelSuccess() {
	s.mockRepo.
		On("GetRandomParagraph", s.ctx).
		Return(s.testParagraph, nil).
		Once()

	respChan := s.builder.GetAsycChannel(s.ctx)
	resp := <-respChan

	assert.NoError(s.T(), resp.Error)
	assert.Equal(s.T(), s.testParagraph, resp.Paragraph)
	s.mockRepo.AssertExpectations(s.T())
}

func (s *ParagraphsBuilderSuite) TestGetAsyncChannelError() {
	mockErr := errors.New("repo error")
	s.mockRepo.
		On("GetRandomParagraph", s.ctx).
		Return((*domain.Paragraph)(nil), mockErr).
		Once()

	respChan := s.builder.GetAsycChannel(s.ctx)
	resp := <-respChan

	assert.Error(s.T(), resp.Error)
	assert.Nil(s.T(), resp.Paragraph)
	s.mockRepo.AssertExpectations(s.T())
}

func (s *ParagraphsBuilderSuite) TestGetAsyncChannelNoParagraphs() {
	s.mockRepo.
		On("GetRandomParagraph", s.ctx).
		Return((*domain.Paragraph)(nil), sql.ErrNoRows).
		Once()
	respChan := s.builder.GetAsycChannel(s.ctx)
	resp := <-respChan
	assert.Equal(s.T(), svcerrors.NoParagraphsFoundError, resp.Error)
	assert.Nil(s.T(), resp.Paragraph)
}

func (s *ParagraphsBuilderSuite) TestBuildSuccess() {
	resp := ParagraphResponseAsync{
		Paragraph: s.testParagraph,
		Error:     nil,
	}
	paragraph, err := s.builder.Build(context.Background(), resp)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), s.testParagraph, paragraph)
}

func (s *ParagraphsBuilderSuite) TestBuildError() {
	testErr := errors.New("some error")
	resp := ParagraphResponseAsync{
		Paragraph: nil,
		Error:     testErr,
	}
	paragraph, err := s.builder.Build(context.Background(), resp)
	assert.Error(s.T(), err)
	assert.Nil(s.T(), paragraph)
}

func TestParagraphsBuilderSuite(t *testing.T) {
	suite.Run(t, new(ParagraphsBuilderSuite))
}
