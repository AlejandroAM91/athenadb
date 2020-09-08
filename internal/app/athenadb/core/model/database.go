package model

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

// AddTable a new table with given name.
func (database *Database) AddTable(name string, table *Table) {
	database.tableMap[name] = table
}

// GetTable gets and returns the table selected by name.
func (database Database) GetTable(name string) (*Table, bool) {
	db, present := database.tableMap[name]
	return db, present
}
