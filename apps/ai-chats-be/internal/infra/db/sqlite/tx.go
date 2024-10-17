package sqlite

import (
	"context"
	"database/sql"
	"fmt"
)

type txKeyType string

const txKey = txKeyType("tx")

func WithTx(ctx context.Context, tx *sql.Tx) context.Context {
	return context.WithValue(ctx, txKey, tx)
}

// MaybeHaveTx returns the transaction from the context if it exists, otherwise returns nil.
func MaybeHaveTx(ctx context.Context) DBTX {
	tx, ok := ctx.Value(txKey).(*sql.Tx)
	if !ok {
		return nil
	}
	return tx
}

// MustHaveTx checks if the transaction exists in the context. If not, it panics.
func MustHaveTx(ctx context.Context) {
	tx := MaybeHaveTx(ctx)
	if tx == nil {
		panic("transaction not found")
	}
}

type Tx struct {
	db *sql.DB
}

func NewTx(db *sql.DB) *Tx {
	return &Tx{db: db}
}

func (t *Tx) Tx(ctx context.Context, fn func(context.Context) error) error {
	tx, err := t.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin a transaction: %w", err)
	}

	ctx = WithTx(ctx, tx)
	if err := fn(ctx); err != nil {
		if err := tx.Rollback(); err != nil {
			return fmt.Errorf("failed to rollback a transaction: %w", err)
		}

		return fmt.Errorf("transaction failed: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit a transaction: %w", err)
	}

	return nil
}
