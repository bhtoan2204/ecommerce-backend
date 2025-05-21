package repository

import (
	"context"
	"user/app/infrastructure/persistent/postgresql/model"

	"gorm.io/gorm"
)

var _ SessionRepository = (*sessionRepository)(nil)

type SessionRepository interface {
	Create(ctx context.Context, session *model.Session) (*model.Session, error)
}

type sessionRepository struct {
	db *gorm.DB
}

func newSessionRepository(db *gorm.DB) SessionRepository {
	return &sessionRepository{
		db: db,
	}
}

func (s *sessionRepository) Create(ctx context.Context, session *model.Session) (*model.Session, error) {
	panic("unimplemented")
}
