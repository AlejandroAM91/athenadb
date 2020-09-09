package services

import (
	"github.com/AlejandroAM91/athenadb/internal/app/athenadb/infra/storage"
)

// CreateDatabase creates a new database
func CreateDatabase(name string) error {
	// TODO: Check database not exists
	return storage.CreateDatabase(name)
}

// RetrieveDatabaseList retreave a list of databases
func RetrieveDatabaseList() ([]string, error) {
	dbList, error := storage.RetrieveDatabaseList()
	return dbList, error
}
