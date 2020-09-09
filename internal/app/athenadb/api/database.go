package api

import (
	"log"
	"net/http"

	"github.com/AlejandroAM91/athenadb/internal/app/athenadb/core/model"
	"github.com/gorilla/mux"

	"github.com/AlejandroAM91/athenadb/internal/app/athenadb/core/ports"
	"github.com/AlejandroAM91/athenadb/internal/pkg/dimgr"
)

// DatabaseController creates database call controller
type DatabaseController struct {
}

// CreateDatabaseController creates table call handler
func CreateDatabaseController() *DatabaseController {
	return &DatabaseController{}
}

// ServeHTTP database controller calls
func (ctrl DatabaseController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		ctrl.post(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
	}
}

func (ctrl DatabaseController) post(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	dimgr := dimgr.GetManager()
	createDbPort := dimgr.GetDependency("CreateDatabasePort").(ports.CreateDatabasePort)
	error := createDbPort.CreateDatabase(vars["database"], model.CreateDatabase())
	if error != nil {
		log.Print(error)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
