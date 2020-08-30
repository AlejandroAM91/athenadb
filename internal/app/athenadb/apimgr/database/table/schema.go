package table

import (
	"net/http"

	"github.com/AlejandroAM91/athenadb/internal/app/athenadb/datamgr"
)

// The SchemaHandler contains a table schema call handler
type SchemaHandler struct {
	datamgr *datamgr.DataManager
}

// CreateSchemaHandler creates table call handler
func CreateSchemaHandler(datamgr *datamgr.DataManager) Handler {
	return Handler{datamgr: datamgr}
}

// ServeHTTP manages schema calls
func (handler SchemaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// database := handler.datamgr.GetDatabase(vars["database"])
	// if database == nil {
	// 	w.WriteHeader(http.StatusNotFound)
	// 	w.Write([]byte(http.StatusText(http.StatusNotFound)))
	// 	return
	// }

	// switch r.Method {
	// case http.MethodGet:
	// 	// schema.get(w, r)
	// default:
	// 	w.WriteHeader(http.StatusNotImplemented)
	// 	w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
	// }
}

// SchemaHandler manages table calls
// func SchemaHandler(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	log.Printf("Called /%s/%s/schema (table schema)\n", vars["database"], vars["table"])
// }
