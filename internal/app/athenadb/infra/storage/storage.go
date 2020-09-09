package storage

import (
    "log"
    "os"
)

// CreateDatabase creates database storage
func CreateDatabase(name string) error {
    _, error := OpenDbFile("admin", os.O_RDWR)
    if error != nil {
        log.Println("[error] CreateDatabase: ", error)
        return error
    }
    return nil
}

// RetrieveDatabaseList retrieve a database list from admin database file
func RetrieveDatabaseList() ([]string, error) {
    _, error := OpenDbFile("admin", os.O_RDONLY)
    if error != nil {
        log.Println("[error] RetrieveDatabaseList: ", error)
        return nil, error
    }
    return []string{}, nil
}
