package ports

import "github.com/AlejandroAM91/athenadb/pkg/core/model"

// Storage represents physical storage manipulation service
type Storage interface {
	GetAll() ([]model.DB, error)
}