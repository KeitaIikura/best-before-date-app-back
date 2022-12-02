package gateway

import (
	"bbdate/internal/bbdate/domain/repository"
	"bbdate/pkg/db"
	"bbdate/pkg/logging"
	"context"
	"database/sql"
	"fmt"
)

type TransactionOnRDB struct {
	Reader *sql.DB
	Writer *sql.DB
}

func NewTransactionOnRDB(m db.IMySQL) repository.Transaction {
	return &TransactionOnRDB{
		Reader: m.GetReaderConn(),
		Writer: m.GetWriterConn(),
	}
}

func (t TransactionOnRDB) Begin(ctx context.Context, xrid string) (*repository.TxWrapper, error) {
	tx, err := t.Writer.BeginTx(ctx, nil)
	if err != nil {
		logging.Error(xrid, fmt.Sprintf("Transaction.Begin Error: %s", err.Error()))
		return nil, err
	}
	return &repository.TxWrapper{
		Tx: tx,
	}, err
}

func (t TransactionOnRDB) Commit(xrid string, txWrapper *repository.TxWrapper) error {
	err := txWrapper.Tx.Commit()
	if err != nil {
		logging.Error(xrid, fmt.Sprintf("Transaction.Commit Error: %s", err.Error()))
		// ロールバック時のエラーは持ち越さない
		_ = t.RollBack(xrid, txWrapper)
		return err
	}
	return nil
}

// 基本的に明示的にロールバックする場合のみ呼び出す想定
func (t TransactionOnRDB) RollBack(xrid string, txWrapper *repository.TxWrapper) error {
	err := txWrapper.Tx.Rollback()
	if err != nil {
		logging.Error(xrid, fmt.Sprintf("Transaction.RollBack Error: %s", err.Error()))
		return err
	}
	logging.Info(xrid, "Transaction.RollBack rollback was successed")
	return nil
}
