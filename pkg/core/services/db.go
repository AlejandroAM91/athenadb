package services

import "github.com/AlejandroAM91/athenadb/pkg/core/model"

// DB represents a database manipulation service
type DB interface {
	// GetAll retrieves a list of databases
	GetAll() ([]model.DB, error)
}
