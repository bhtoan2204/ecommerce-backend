package entities

import (
	"time"
)

type Session struct {
	userID       int64
	refreshToken string
	userAgent    string
	ipAddress    string
	isRevoked    bool
	expiresAt    time.Time
}

func NewSession(userID int64, refreshToken, userAgent, ipAddress string, isRevoked bool, expiresAt time.Time) *Session {
	return &Session{
		userID:       userID,
		refreshToken: refreshToken,
		userAgent:    userAgent,
		ipAddress:    ipAddress,
		isRevoked:    isRevoked,
		expiresAt:    expiresAt,
	}
}

func (s *Session) UserID() int64 {
	return s.userID
}

func (s *Session) RefreshToken() string {
	return s.refreshToken
}

func (s *Session) UserAgent() string {
	return s.userAgent
}

func (s *Session) IPAddress() string {
	return s.ipAddress
}

func (s *Session) IsRevoked() bool {
	return s.isRevoked
}

func (s *Session) ExpiresAt() time.Time {
	return s.expiresAt
}

func (s *Session) SetExpiresAt(expiresAt time.Time) {
	s.expiresAt = expiresAt
}

func (s *Session) SetIsRevoked(isRevoked bool) {
	s.isRevoked = isRevoked
}

func (s *Session) SetRefreshToken(refreshToken string) {
	s.refreshToken = refreshToken
}

func (s *Session) SetUserAgent(userAgent string) {
	s.userAgent = userAgent
}

func (s *Session) SetIPAddress(ipAddress string) {
	s.ipAddress = ipAddress
}

func (s *Session) SetUserID(userID int64) {
	s.userID = userID
}
