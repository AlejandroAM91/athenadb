package schema

import (
	"net/http"

	"github.com/AlejandroAM91/athenadb/internal/app/athenadb/datamgr"

	"github.com/gorilla/mux"
)

// The Handler contains a database schema call handler
type Handler struct {
	datamgr *datamgr.DataManager
}

// CreateHandler creates table call handler
func CreateHandler(datamgr *datamgr.DataManager) Handler {
	return Handler{datamgr: datamgr}
}

// ServeHTTP manages schema calls
func (handler Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_, present := handler.datamgr.GetDatabase(vars["database"])
	if !present {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(http.StatusText(http.StatusNotFound)))
		return
	}

	// switch r.Method {
	// case http.MethodGet:
	// schema.get(w, r)
	// default:
	// 	w.WriteHeader(http.StatusNotImplemented)
	// 	w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
	// }
}
