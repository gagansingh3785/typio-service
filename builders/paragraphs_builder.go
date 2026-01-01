package builders

import (
	"context"
	"database/sql"
	"errors"

	"github.com/gagansingh3785/typio-service/domain"
	svcerrors "github.com/gagansingh3785/typio-service/errors"
	"github.com/gagansingh3785/typio-service/repository"
)

type ParagraphResponseAsync struct {
	Paragraph *domain.Paragraph
	Error     error
}

type ParagraphsBuilder interface {
	GetAsycChannel(ctx context.Context) chan ParagraphResponseAsync
	Build(ctx context.Context, resp ParagraphResponseAsync) (*domain.Paragraph, error)
}

type paragraphsBuilder struct {
	paragraphsRepo repository.Repository
}

func NewParagraphsBuilder(paragraphsRepo repository.Repository) ParagraphsBuilder {
	return &paragraphsBuilder{
		paragraphsRepo: paragraphsRepo,
	}
}

func (b *paragraphsBuilder) GetAsycChannel(ctx context.Context) chan ParagraphResponseAsync {
	respChan := make(chan ParagraphResponseAsync, 1)
	go func() {
		defer close(respChan)
		paragraph, err := b.paragraphsRepo.GetRandomParagraph(ctx)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				err = svcerrors.NoParagraphsFoundError
			}
			respChan <- ParagraphResponseAsync{
				Paragraph: nil,
				Error:     err,
			}
			return
		}

		respChan <- ParagraphResponseAsync{
			Paragraph: paragraph,
			Error:     nil,
		}
	}()

	return respChan
}

func (b *paragraphsBuilder) Build(ctx context.Context, resp ParagraphResponseAsync) (*domain.Paragraph, error) {
	if resp.Error != nil {
		return nil, resp.Error
	}
	return resp.Paragraph, nil
}
