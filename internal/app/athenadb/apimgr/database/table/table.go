package table

import (
	"net/http"

	"github.com/AlejandroAM91/athenadb/internal/app/athenadb/datamgr"
)

// The Handler contains a table call handler
type Handler struct {
	datamgr *datamgr.DataManager
}

// CreateHandler creates table call handler
func CreateHandler(datamgr *datamgr.DataManager) Handler {
	return Handler{datamgr: datamgr}
}

// ServeHTTP manages table calls
func (handler Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// database := schema.datamgr.GetDatabase(vars["database"])
	// if database == nil {
	// 	w.WriteHeader(http.StatusNotFound)
	// 	w.Write([]byte(http.StatusText(http.StatusNotFound)))
	// 	return
	// }

	// switch r.Method {
	// default:
	// 	w.WriteHeader(http.StatusNotImplemented)
	// 	w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
	// }
}

// Handler manages table calls
// func Handler(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	log.Printf("Called /%s/%s (table data)\n", vars["database"], vars["table"])
// }
