package services

import "github.com/AlejandroAM91/athenadb/pkg/domain/model"

// DB represents a database manipulation service
type DB interface {
	GetAll() ([]model.DB, error)
}
