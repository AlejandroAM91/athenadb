package datamgr

import "github.com/AlejandroAM91/athenadb/internal/app/athenadb/core/model"

const (
	databaseAdmin = "admin"
	tableSchema   = "schema"
)

// The DataManager manages application data operations (load, save, queries...).
type DataManager struct {
	databaseMap map[string]*model.Database
}

// CreateDataManager creates and initializes aplication data manager.
func CreateDataManager() *DataManager {
	databaseMap := make(map[string]*model.Database)

	// Set default values
	databaseMap[databaseAdmin] = model.CreateDatabase()

	return &DataManager{
		databaseMap: databaseMap,
	}
}

// GetDatabase gets and returns the database selected by name.
func (datamgr DataManager) GetDatabase(name string) (*model.Database, bool) {
	db, present := datamgr.databaseMap[name]
	return db, present
}
