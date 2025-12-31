package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type TxRunner interface {
	RunInTxContext(ctx context.Context, fn func(*sqlx.Tx) error) error
}

type txRunner struct {
	db *sqlx.DB
}

func NewTxRunner(db *sqlx.DB) TxRunner {
	return &txRunner{
		db: db,
	}
}

func (r *txRunner) RunInTxContext(ctx context.Context, fn func(*sqlx.Tx) error) error {
	// start a transaction
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	// stop panic of program
	defer func() {
		if ret := recover(); ret != nil {
			// there was a panic, need to roll back the transaction
			_ = tx.Rollback()
			err = errors.New(fmt.Sprint(ret))
		}
	}()

	// run the function in the transaction
	if err = fn(tx); err != nil {
		_ = tx.Rollback()
		return err
	}

	// commit for successful transaction
	return tx.Commit()
}
