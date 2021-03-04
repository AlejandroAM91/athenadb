package web

import (
	"encoding/json"
	"net/http"

	"github.com/AlejandroAM91/athenadb/internal/app/server/core/services"
)

// DBCtrl represents an database http controller
type DBCtrl interface {
	GetAll(w http.ResponseWriter, r *http.Request)
}

type dbcrtl struct {
	dbsrv services.DB
}

// NewDBCtrl creates a database http controller
func NewDBCtrl(dbsrv services.DB) DBCtrl {
	return &dbcrtl{dbsrv: dbsrv}
}

// GetAll retrieves a list of databases
func (ctrl dbcrtl) GetAll(w http.ResponseWriter, r *http.Request) {
	dblist, error := ctrl.dbsrv.GetAll()
	if error != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Process output
	error = json.NewEncoder(w).Encode(dblist)
	if error != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
