package repository

import (
	"context"
	"payment/package/logger"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository interface {
	WithTransaction(ctx context.Context, fn func(Repository) error) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {

	return &repository{
		db: db,
	}
}

// WithTransaction ...
func (r *repository) WithTransaction(ctx context.Context, fn func(Repository) error) (err error) {
	log := logger.FromContext(ctx)
	log.Info("Starting transaction")

	tx := r.db.Begin()
	tr := NewRepository(tx)

	err = tx.Error
	if err != nil {
		return
	}

	defer func() {
		if p := recover(); p != nil { // nolint
			log.Error("Transaction failed with panic: ", zap.Any("panic", p))
			// a panic occurred, rollback and repanic
			tx.Rollback()
			panic(p)
		} else if err != nil {
			log.Warn("Transaction failure with error: ", zap.Error(err))
			// something went wrong, rollback
			tx.Rollback()
		} else {
			log.Info("Finishing transaction")
			// all good, commit
			err = tx.Commit().Error
		}
	}()

	err = fn(tr)

	return err
}
