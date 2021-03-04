package services

import (
	"errors"

	"github.com/AlejandroAM91/athenadb/internal/app/server/core/model"
	"github.com/AlejandroAM91/athenadb/internal/app/server/ports"
)

// DB represents a database manipulation service
type DB interface {
	// GetAll retrieves a list of databases
	GetAll() ([]model.DB, error)
}

type dbsrv struct {
	storage ports.Storage
}

// NewDBSrv creates a database manipulation service
func NewDBSrv(storage ports.Storage) DB {
	return &dbsrv{storage: storage}
}

// GetAll retrieves a list of databases
func (s dbsrv) GetAll() ([]model.DB, error) {
	dblist, err := s.storage.GetAll()
	if err != nil {
		return []model.DB{}, errors.New("Retrieve database list from storage failed")
	}
	return dblist, nil
}