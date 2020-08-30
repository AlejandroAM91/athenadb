package datamgr

// The DataManager manages application data operations (load, save, queries...).
type DataManager struct {
	DatabaseList []database
}

// CreateDataManager creates and initializes aplication data manager.
func CreateDataManager() *DataManager {
	databaseList := make([]database, 0)

	// Set default values
	databaseList = append(databaseList, database{Name: "admin"})

	return &DataManager{
		DatabaseList: databaseList,
	}
}

type database struct {
	Name string
	// TableList []table
}

type table struct {
	Name string
}
