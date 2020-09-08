package model

// The Database contains the information about a database
type Database struct{}

// CreateDatabase creates and initializes a database.
func CreateDatabase() *Database {
	return &Database{}
}
