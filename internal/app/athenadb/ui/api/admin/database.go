package admin

import (
	"encoding/json"
        "log"
	"net/http"

	"github.com/AlejandroAM91/athenadb/internal/app/athenadb/core/services"
)

// DatabaseController creates database call controller
type DatabaseController struct{}

// CreateDatabaseController creates table call handler
func CreateDatabaseController() DatabaseController {
	return DatabaseController{}
}

// ServeHTTP database controller calls
func (ctrl DatabaseController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		ctrl.get(w, r)
	case http.MethodPost:
		ctrl.post(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
	}
}

func (ctrl DatabaseController) get(w http.ResponseWriter, r *http.Request) {
	// Process request
	dbList, error := services.RetrieveDatabaseList()
	if error != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Process output
	error = json.NewEncoder(w).Encode(dbList)
	if error != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (ctrl DatabaseController) post(w http.ResponseWriter, r *http.Request) {
	var input struct{ Name string }

	// Process input
	error := json.NewDecoder(r.Body).Decode(&input)
	if error != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Process request
	error = services.CreateDatabase(input.Name)
	if error != nil {
                log.Println(error)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Process output
	w.WriteHeader(http.StatusCreated)
}
