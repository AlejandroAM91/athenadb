package ports

import (
	"github.com/AlejandroAM91/athenadb/internal/app/server/core/model"
)

// Storage represents physical storage manipulation service
type Storage interface {
	// GetAll retrieves a list of databases
	GetAll() ([]model.DB, error)
}

type storage struct {
}

// NewStorage creates storage manipulation adapter
func NewStorage() Storage {
	return &storage{}
}

// GetAll retrieves a list of databases
func (s storage) GetAll() ([]model.DB, error) {
	return make([]model.DB, 0), nil
}
