package store

import (
	"errors" // constructs error values
	"strconv"
	"time" // timestamp generation and formatting
)

type User struct {
	Username  string
	Password  string
	AuthToken string
}

// --- CUSTOMER MODEL ---
// Data struct with JSON tags.
type Customer struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Comment   string `json:"comment"`
	Status    string `json:"status"`
	Score     int    `json:"score"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type MemoryStore struct {
	customers map[string]*Customer
	users     map[string]*User
	nextID    int
}

// Allocates MemoryStore and initializes maps with data.
func NewMemoryStore() *MemoryStore {

	// --- TIMESTAMP CREATION ---
	// time.Now()      → current local time
	// .UTC()          → convert to UTC timezone
	// .Format(...)    → string representation
	now := time.Now().UTC().Format(time.DateTime)

	// &MemoryStore{...} allocates on heap.
	return &MemoryStore{
		customers: map[string]*Customer{
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
		users: map[string]*User{
			"admin": {
				Username:  "admin",
				Password:  "admin",
				AuthToken: "demo-token",
			},
		},
		nextID: 2,
	}
}

func (s *MemoryStore) GetUser(username string) *User {
	user, ok := s.users[username]
	if !ok {
		return nil
	}
	return user
}

func (s *MemoryStore) GetAllCustomers() []*Customer {

	// result is nil slice initially.
	var result []*Customer

	// gets all customer from ID 1 until max ID (nextID)
	for i := 1; i < s.nextID; i++ {
		if c, ok := s.customers[strconv.Itoa(i)]; ok {
			result = append(result, c)
		}
	}
	return result
}

func (s *MemoryStore) GetCustomer(id string) (*Customer, error) {

	customer, ok := s.customers[id]
	if !ok {
		return nil, errors.New("customer doesn't exist")
	}

	return customer, nil
}

func (s *MemoryStore) CreateCustomer(c *Customer) error {
	now := time.Now().UTC().Format(time.DateTime)
	_, exists := s.customers[c.ID]
	if exists {
		return errors.New("already exists")
	}
	id := strconv.Itoa(s.nextID)
	s.customers[id] = &Customer{
		ID:        id,
		Name:      c.Name,
		Comment:   c.Comment,
		Status:    c.Status,
		Score:     c.Score,
		CreatedAt: now,
		UpdatedAt: now,
	}
	s.nextID++
	return nil
}

func (s *MemoryStore) UpdateCustomer(id string, c *Customer) bool {
	_, exists := s.customers[id]
	if exists {
		s.customers[id].Name = c.Name
		s.customers[id].Comment = c.Comment
		s.customers[id].Score = c.Score
		s.customers[id].Status = c.Status
		s.customers[id].UpdatedAt = time.Now().UTC().Format(time.DateTime)
		return true
	}
	return false
}

func (s *MemoryStore) DeleteCustomer(id string) bool {
	_, exists := s.customers[id]
	if exists {
		delete(s.customers, id)
		return true
	}
	return false
}
