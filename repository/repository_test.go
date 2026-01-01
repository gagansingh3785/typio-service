package repository

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/gagansingh3785/typio-service/appcontext"
	"github.com/gagansingh3785/typio-service/config"
	"github.com/gagansingh3785/typio-service/database"
	"github.com/gagansingh3785/typio-service/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RepositoryTestSuite struct {
	suite.Suite

	ctx        context.Context
	repository Repository
	db         *database.Database
	paragraphs []domain.Paragraph
}

func TestRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}

func (s *RepositoryTestSuite) SetupSuite() {
	cfg, err := config.SetupConfig("test_application")
	if err != nil {
		s.T().Fatalf("Failed to setup config: %v", err)
	}

	appcontext.Initiate(cfg)
	s.ctx = context.Background()
	s.db = appcontext.GetDatabase()
	s.repository = NewRepository(s.db)
	now := time.Now()
	s.paragraphs = []domain.Paragraph{
		{
			Content:   "The quick brown fox jumps over the lazy dog.",
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			Content:   "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			Content:   "Pack my box with five dozen liquor jugs.",
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			Content:   "Sphinx of black quartz, judge my vow.",
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			Content:   "How razorback-jumping frogs can level six piqued gymnasts!",
			CreatedAt: now,
			UpdatedAt: now,
		},
	}
}

func (s *RepositoryTestSuite) SetupTest() {
	for i, paragraph := range s.paragraphs {
		insertedParagraph, err := s.repository.InsertParagraph(s.ctx, paragraph)
		assert.NoError(s.T(), err)
		s.paragraphs[i] = *insertedParagraph
	}
}

func (s *RepositoryTestSuite) TearDownTest() {
	_, err := s.db.DB.Exec("DELETE FROM paragraphs")
	assert.NoError(s.T(), err)
}

func (s *RepositoryTestSuite) TestGetRandomParagraph() {
	paragraph, err := s.repository.GetRandomParagraph(s.ctx)
	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), paragraph)
	assert.Contains(s.T(), s.paragraphs, *paragraph)
}

func (s *RepositoryTestSuite) TestGetRandomParagraphNoParagraphs() {
	_, err := s.db.DB.Exec("DELETE FROM paragraphs")
	assert.NoError(s.T(), err)
	paragraph, err := s.repository.GetRandomParagraph(s.ctx)
	assert.Equal(s.T(), sql.ErrNoRows, err)
	assert.Nil(s.T(), paragraph)
}
