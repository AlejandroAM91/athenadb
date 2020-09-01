package table

import (
	"net/http"

	"github.com/AlejandroAM91/athenadb/internal/app/athenadb/datamgr"
	"github.com/gorilla/mux"
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
	vars := mux.Vars(r)
	db, present := handler.datamgr.GetDatabase(vars["database"])
	if !present {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(http.StatusText(http.StatusNotFound)))
		return
	}

	_, present = db.GetTable(vars["table"])
	if !present {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(http.StatusText(http.StatusNotFound)))
		return
	}

	// switch r.Method {
	// case http.MethodGet:
	// 	// schema.get(w, r)
	// default:
	// 	w.WriteHeader(http.StatusNotImplemented)
	// 	w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
	// }
}
