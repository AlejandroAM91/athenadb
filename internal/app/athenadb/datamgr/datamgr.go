package datamgr

const (
	databaseAdmin = "admin"
	tableSchema   = "schema"
)

// The DataManager manages application data operations (load, save, queries...).
type DataManager struct {
	databaseMap map[string]*Database
}

// CreateDataManager creates and initializes aplication data manager.
func CreateDataManager() *DataManager {
	databaseMap := make(map[string]*Database)

	// Set default values
	databaseMap[databaseAdmin] = CreateDatabase()

	return &DataManager{
		databaseMap: databaseMap,
	}
}

// GetDatabase gets and returns the database selected by name.
func (datamgr DataManager) GetDatabase(name string) (*Database, bool) {
	db, present := datamgr.databaseMap[name]
	return db, present
}
