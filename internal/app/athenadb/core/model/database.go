package model

// The Database contains the information about a database
type Database struct {
	colList []string
}

// CreateDatabase creates and initializes a database.
func CreateDatabase(colList []string) *Database {
	return &Database{
		colList: colList,
	}
}
