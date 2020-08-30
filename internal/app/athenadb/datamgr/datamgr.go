package datamgr

// The DataManager manages application data operations (load, save, queries...).
type DataManager struct {
	databaseList []Database
}

// CreateDataManager creates and initializes aplication data manager.
func CreateDataManager() *DataManager {
	databaseList := make([]Database, 0)

	// Set default values
	databaseList = append(databaseList, Database{Name: "admin"})

	return &DataManager{
		databaseList: databaseList,
	}
}

// GetDatabase gets and returns the database selected by name.
func (datamgr DataManager) GetDatabase(name string) *Database {
	databaseList := datamgr.databaseList
	for i := 0; i < len(databaseList); i++ {
		if databaseList[i].Name == name {
			return &databaseList[i]
		}
	}
	return nil
}
