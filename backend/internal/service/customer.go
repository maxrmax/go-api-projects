package service

import (
	"backend/internal/store" // provides MemoryStore and Customer
	"errors"
)

// This is a pointer reference to the data layer.
// No customer data is stored inside this struct.
type CustomerService struct {
	store *store.MemoryStore
}

func NewCustomerService(s *store.MemoryStore) *CustomerService {
	return &CustomerService{store: s}
}

func (s *CustomerService) GetAllCustomers() []*store.Customer {
	return s.store.GetAllCustomers()
}

// get a single Customer
func (s *CustomerService) GetCustomer(id string) (*store.Customer, error) {
	return s.store.GetCustomer(id)
}

// Create a Customer
func (s *CustomerService) CreateCustomer(c *store.Customer) error {
	err := s.store.CreateCustomer(c)
	if err != nil {
		return errors.New("couldn't create")
	}
	return nil
}

func (s *CustomerService) UpdateCustomer(id string, c *store.Customer) error {
	ok := s.store.UpdateCustomer(id, c)
	if !ok {
		return errors.New("customer not found")
	}
	return nil
}

func (s *CustomerService) DeleteCustomer(id string) error {
	ok := s.store.DeleteCustomer(id)
	if !ok {
		return errors.New("customer not found")
	}
	return nil
}
