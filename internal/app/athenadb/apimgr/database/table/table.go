package table

import (
	"net/http"

	"github.com/AlejandroAM91/athenadb/internal/app/athenadb/datamgr"
	"github.com/gorilla/mux"
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
	// default:
	// 	w.WriteHeader(http.StatusNotImplemented)
	// 	w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
	// }
}
