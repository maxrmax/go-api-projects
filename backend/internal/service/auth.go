package service

import "backend/internal/store" // provides MemoryStore and User

type AuthService struct {
	store *store.MemoryStore
}

func NewAuthService(s *store.MemoryStore) *AuthService {
	return &AuthService{store: s}
}

// demo of auth implementation
// purely for demonstative aspects
func (s *AuthService) Login(username, password string) string {

	user := s.store.GetUser(username)
	if user == nil {
		return ""
	}

	if user.Password != password {
		return ""
	}
	return user.AuthToken
}
