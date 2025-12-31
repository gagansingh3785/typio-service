package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type BaseRepository interface {
	NamedGetRow(ctx context.Context, tx *sqlx.Tx, query string, dest interface{}, arg interface{}) error
	NamedQueryRow(ctx context.Context, tx *sqlx.Tx, query string, dest interface{}, arg interface{}) error
}

type baseRepository struct {
	db *sqlx.DB
}

func NewBaseRepository(db *sqlx.DB) BaseRepository {
	return &baseRepository{
		db: db,
	}
}

func (r *baseRepository) NamedGetRow(ctx context.Context, tx *sqlx.Tx, query string, dest interface{}, arg interface{}) error {
	query, args, err := r.db.BindNamed(query, arg)
	if err != nil {
		return err
	}

	if tx != nil {
		return tx.GetContext(ctx, dest, query, args...)
	}

	return r.db.GetContext(ctx, dest, query, args...)
}

func (r *baseRepository) NamedQueryRow(ctx context.Context, tx *sqlx.Tx, query string, dest interface{}, arg interface{}) error {
	query, args, err := r.db.BindNamed(query, arg)
	if err != nil {
		return err
	}

	if tx != nil {
		row := tx.QueryRowxContext(ctx, query, args...)
		return row.StructScan(dest)
	}

	row := r.db.QueryRowxContext(ctx, query, args...)
	return row.StructScan(dest)
}
