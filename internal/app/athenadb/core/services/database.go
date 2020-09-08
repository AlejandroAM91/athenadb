package services

import (
	"errors"

	"github.com/AlejandroAM91/athenadb/internal/app/athenadb/core/model"
)

// The DatabaseService contains the information about all databases
type DatabaseService struct {
	databaseMap map[string]*model.Database
}

// CreateDatabaseService creates and initializes the database service.
func CreateDatabaseService() *DatabaseService {
	return &DatabaseService{
		databaseMap: make(map[string]*model.Database),
	}
}

// CreateDatabase creates a database.
func (dbsrv *DatabaseService) CreateDatabase(name string, database *model.Database) error {
	_, exists := dbsrv.databaseMap[name]
	if exists {
		return errors.New("Database already exists")
	}
	dbsrv.databaseMap[name] = database
	return nil
}

// RetrieveDatabase retrieves database schema.
func (dbsrv DatabaseService) RetrieveDatabase(name string) (*model.Database, error) {
	db, exists := dbsrv.databaseMap[name]
	if !exists {
		return nil, errors.New("Database not exists")
	}
	return db, nil
}

// UpdateDatabase creates and initializes a database.
func (dbsrv *DatabaseService) UpdateDatabase(name string, database *model.Database) error {
	_, exists := dbsrv.databaseMap[name]
	if !exists {
		return errors.New("Database not exists")
	}
	dbsrv.databaseMap[name] = database
	return nil
}

// DeleteDatabase deletes a database.
func (dbsrv *DatabaseService) DeleteDatabase(name string) error {
	_, exists := dbsrv.databaseMap[name]
	if !exists {
		return errors.New("Database not exists")
	}
	delete(dbsrv.databaseMap, name)
	return nil
}
