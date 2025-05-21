package mapper

import (
	"user/app/domain/entities"
	"user/app/infrastructure/persistent/postgresql/model"
)

func SessionToModel(session *entities.Session) *model.Session {
	return &model.Session{
		UserID:       session.UserID(),
		RefreshToken: session.RefreshToken(),
		UserAgent:    session.UserAgent(),
		IPAddress:    session.IPAddress(),
		IsRevoked:    session.IsRevoked(),
		ExpiresAt:    session.ExpiresAt(),
	}
}

func SessionToEntity(session *model.Session) *entities.Session {
	return entities.NewSession(
		session.UserID,
		session.RefreshToken,
		session.UserAgent,
		session.IPAddress,
		session.IsRevoked,
		session.ExpiresAt,
	)
}
