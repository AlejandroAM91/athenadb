package ports

import "github.com/AlejandroAM91/athenadb/pkg/domain/model"

// Storage represents physical storage manipulation service
type Storage interface {
	GetAll() ([]model.DB, error)
}
