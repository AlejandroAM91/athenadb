package api

import (
	"net/http"

	"github.com/AlejandroAM91/athenadb/internal/app/athenadb/core/services"
	"github.com/gorilla/mux"
)

// DatabaseController creates database call controller
type DatabaseController struct {
	services *services.Services
}

// CreateDatabaseController creates table call handler
func CreateDatabaseController(services *services.Services) *DatabaseController {
	return &DatabaseController{
		services: services,
	}
}

// ServeHTTP database controller calls
func (ctrl DatabaseController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		ctrl.get(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
	}
}

func (ctrl DatabaseController) get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_, error := ctrl.services.GetDatabaseService().RetrieveDatabase(vars["database"])
	if error != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
}
