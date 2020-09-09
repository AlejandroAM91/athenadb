package storage

import "os"

// The DbFile contaions database file information
type DbFile struct{}

// OpenDbFile opens database file with given database name
func OpenDbFile(dbname string, flag int) (*DbFile, error) {
    path := "db/base/"+dbname+".adb"

    _, error := os.OpenFile(path, flag | os.O_CREATE, 0600)
    if error != nil {
        return nil, error
    }
    return &DbFile{}, nil
}
