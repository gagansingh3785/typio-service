package services

import (
	"context"

	"github.com/gagansingh3785/typio-service/builders"
	"github.com/gagansingh3785/typio-service/domain"
)

type ParagraphService interface {
	GetRandomParagraph(ctx context.Context) (*domain.Paragraph, error)
}

type paragraphService struct {
	paragraphsBuilder builders.ParagraphsBuilder
}

func NewParagraphService(paragraphsBuilder builders.ParagraphsBuilder) ParagraphService {
	return &paragraphService{
		paragraphsBuilder: paragraphsBuilder,
	}
}

func (s *paragraphService) GetRandomParagraph(ctx context.Context) (*domain.Paragraph, error) {
	respChan := s.paragraphsBuilder.GetAsycChannel(ctx)

	return s.paragraphsBuilder.Build(ctx, <-respChan)
}
