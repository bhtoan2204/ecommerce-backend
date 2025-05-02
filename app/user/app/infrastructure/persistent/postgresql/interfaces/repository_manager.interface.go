package interfaces

import "context"

type RepositoryManager interface {
	WithTransaction(ctx context.Context, fn func(RepositoryManager) error) error
	// UserRepository() UserRepository
}

// type UserRepository interface {
