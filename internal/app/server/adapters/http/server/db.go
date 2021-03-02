package server

import (
	"encoding/json"
	"net/http"

	"github.com/AlejandroAM91/athenadb/pkg/core/services"
	"github.com/go-chi/chi/v5"
)

type dbhdl struct {
	dbsrv services.DB
}

// NewDBHdl creates a database http controller
func NewDBHdl(dbsrv services.DB) http.Handler {
	h := dbhdl{
		dbsrv: dbsrv,
	}
	r := chi.NewRouter()
	r.Get("/", h.getAll)
	return r
}

func (h dbhdl) getAll(w http.ResponseWriter, r *http.Request) {
	dblist, error := h.dbsrv.GetAll()
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
