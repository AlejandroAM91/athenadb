package services

import (
	"errors"

	"github.com/AlejandroAM91/athenadb/internal/app/server/ports"
	"github.com/AlejandroAM91/athenadb/pkg/core/model"
	"github.com/AlejandroAM91/athenadb/pkg/core/services"
)

type dbsrv struct {
	storage ports.Storage
}

// NewDBSrv creates a database manipulation service
func NewDBSrv(storage ports.Storage) services.DB {
	return &dbsrv{
		storage: storage,
	}
}

func (s dbsrv) GetAll() ([]model.DB, error) {
	dblist, err := s.storage.GetAll()
	if err != nil {
		return []model.DB{}, errors.New("Retrieve database list from storage failed")
	}
	return dblist, nil
}
