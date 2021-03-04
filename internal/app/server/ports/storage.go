package ports

import "github.com/AlejandroAM91/athenadb/internal/app/server/core/model"

// Storage represents physical storage manipulation service
type Storage interface {
	// GetAll retrieves a list of databases
	GetAll() ([]model.DB, error)
}
