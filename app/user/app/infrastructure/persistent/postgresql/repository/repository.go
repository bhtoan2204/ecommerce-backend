package repository

import (
	"context"
	"user/package/logger"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository interface {
	WithTransaction(ctx context.Context, fn func(Repository) error) error

	SessionRepository() SessionRepository
	AddressRepository() AddressRepository
	UserRepository() UserRepository
}

type repository struct {
	db *gorm.DB

	sessionRepository SessionRepository
	addressRepository AddressRepository
	userRepository    UserRepository
}

func NewRepository(db *gorm.DB) Repository {
	sessionRepository := newSessionRepository(db)
	addressRepository := newAddressRepository(db)
	userRepository := newUserRepository(db)

	return &repository{
		db:                db,
		sessionRepository: sessionRepository,
		addressRepository: addressRepository,
		userRepository:    userRepository,
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

func (r *repository) SessionRepository() SessionRepository {
	return r.sessionRepository
}

func (r *repository) AddressRepository() AddressRepository {
	return r.addressRepository
}

func (r *repository) UserRepository() UserRepository {
	return r.userRepository
}
