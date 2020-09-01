package datamgr

// The Database contains the information about a database
type Database struct {
	tableMap map[string]*Table
}

// CreateDatabase creates and initializes a database.
func CreateDatabase() *Database {
	return &Database{
		tableMap: make(map[string]*Table),
	}
}

// GetTable gets and returns the table selected by name.
func (database Database) GetTable(name string) (*Table, bool) {
	db, present := database.tableMap[name]
	return db, present
}
