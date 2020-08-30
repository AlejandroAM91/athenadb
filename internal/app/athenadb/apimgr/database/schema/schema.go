package schema

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AlejandroAM91/athenadb/internal/app/athenadb/datamgr"

	"github.com/gorilla/mux"
)

// The Schema contains a database schema information
type Schema struct {
	datamgr *datamgr.DataManager
}

// CreateSchema creates database schema handler
func CreateSchema(datamgr *datamgr.DataManager) Schema {
	return Schema{datamgr: datamgr}
}

// Handler manages schema calls
func (schema Schema) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Printf("Called /%s/schema (database schema)\n", vars["database"])
	databaseList := schema.datamgr.DatabaseList
	pos := -1
	for i := 0; i < len(databaseList); i++ {
		if databaseList[i].Name == vars["database"] {
			pos = i
		}
	}
	if pos < 0 {
		fmt.Println("Database not found")
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Println("Database found")
}
