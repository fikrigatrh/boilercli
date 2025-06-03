package tx

import (
	"boilerplate/config/infra"
	"context"
	"errors"
	"gorm.io/gorm"
)

type TxManager interface {
	Begin(ctx context.Context) (context.Context, error)
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

type txManager struct {
	db *gorm.DB
}

// Create a key type for storing tx in context
type txKey struct{}

// ProvideTxManager creates a new transaction manager
func ProvideTxManager(infra *infra.Infra) TxManager {
	return &txManager{
		db: infra.DbPsql.DB,
	}
}

// Begin starts a new transaction and stores it in context
func (t *txManager) Begin(ctx context.Context) (context.Context, error) {
	tx := t.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return ctx, tx.Error
	}
	return context.WithValue(ctx, txKey{}, tx), nil
}

// Commit commits the transaction stored in context
func (t *txManager) Commit(ctx context.Context) error {
	tx, ok := ctx.Value(txKey{}).(*gorm.DB)
	if !ok {
		return ErrNoTransaction
	}
	return tx.Commit().Error
}

// Rollback aborts the transaction stored in context
func (t *txManager) Rollback(ctx context.Context) error {
	tx, ok := ctx.Value(txKey{}).(*gorm.DB)
	if !ok {
		return ErrNoTransaction
	}
	return tx.Rollback().Error
}

// GetDB extracts DB from context (either transaction or regular DB)
func GetDB(ctx context.Context, defaultDB *gorm.DB) *gorm.DB {
	tx, ok := ctx.Value(txKey{}).(*gorm.DB)
	if ok {
		return tx
	}
	return defaultDB.WithContext(ctx)
}

// ErrNoTransaction is returned when no transaction is found in context
var ErrNoTransaction = errors.New("no transaction in context")
