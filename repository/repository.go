package repository

import (
	"context"
	"crypto/rand"
	"database/sql"
	"math/big"

	"github.com/gagansingh3785/typio-service/database"
	"github.com/gagansingh3785/typio-service/domain"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	InsertParagraph(ctx context.Context, paragraph domain.Paragraph) (*domain.Paragraph, error)
	GetRandomParagraph(ctx context.Context) (*domain.Paragraph, error)
}

type repo struct {
	baseRepository BaseRepository
	txRunner       TxRunner
}

func NewRepository(db *database.Database) Repository {
	return &repo{
		baseRepository: NewBaseRepository(db.DB),
		txRunner:       NewTxRunner(db.DB),
	}
}

func (r *repo) GetRandomParagraph(ctx context.Context) (*domain.Paragraph, error) {
	paragraph := &domain.Paragraph{}
	paragraphCount := 0
	if err := r.txRunner.RunInTxContext(ctx, func(tx *sqlx.Tx) error {
		if err := r.baseRepository.NamedGetRow(ctx, tx, GetParagraphsCountQuery, &paragraphCount, map[string]interface{}{}); err != nil {
			return err
		}

		if paragraphCount == 0 {
			return sql.ErrNoRows
		}

		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(paragraphCount)))
		if err != nil {
			return err
		}

		randomRowNum := randomIndex.Int64() + 1
		if err := r.baseRepository.NamedGetRow(ctx, tx, GetParagraphByIDQuery, paragraph, map[string]interface{}{
			"row_num": randomRowNum,
		}); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return paragraph, nil
}

func (r *repo) InsertParagraph(ctx context.Context, paragraph domain.Paragraph) (*domain.Paragraph, error) {
	insertedParagraph := &domain.Paragraph{}
	if err := r.baseRepository.NamedQueryRow(ctx, nil, InsertParagraphsQuery, insertedParagraph, map[string]interface{}{
		"content":    paragraph.Content,
		"created_at": paragraph.CreatedAt,
		"updated_at": paragraph.UpdatedAt,
	}); err != nil {
		return nil, err
	}

	return insertedParagraph, nil
}
