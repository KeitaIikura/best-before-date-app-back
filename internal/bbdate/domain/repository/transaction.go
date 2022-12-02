package repository

import (
	"context"
	"database/sql"
)

type TxWrapper struct {
	Tx *sql.Tx
}

type Transaction interface {
	Begin(ctx context.Context, xrid string) (*TxWrapper, error)
	Commit(xrid string, tx *TxWrapper) error
	RollBack(xrid string, tx *TxWrapper) error
}
