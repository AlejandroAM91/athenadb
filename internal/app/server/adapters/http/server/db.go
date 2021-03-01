package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type dbhdl struct{}

// NewDBHdl creates a database http controller
func NewDBHdl() http.Handler {
	h := dbhdl{}
	r := chi.NewRouter()
	r.Get("/", h.getAll)
	return r
}

func (h dbhdl) getAll(w http.ResponseWriter, r *http.Request) {
}
