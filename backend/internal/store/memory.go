package store

import (
	"time"

	"go-api-project/backend/internal/model"
)

type MemoryStore struct {
	customers map[string]*model.Customer
	users     map[string]*model.User
}

func NewMemoryStore() *MemoryStore {
	now := time.Now().UTC().Format(time.RFC3339)

	return &MemoryStore{
		customers: map[string]*model.Customer{
			"1": {
				ID:        "1",
				Name:      "Jotaro",
				Comment:   "Stand user",
				Status:    "active",
				Score:     100,
				CreatedAt: now,
				UpdatedAt: now,
			},
		},
		users: map[string]*model.User{
			"admin": {
				Username:  "admin",
				Password:  "login",
				AuthToken: "demo-token",
			},
		},
	}
}
