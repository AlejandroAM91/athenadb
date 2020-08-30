package apimgr

import (
	"log"
	"net/http"
	"os"

	"github.com/AlejandroAM91/athenadb/internal/app/athenadb/apimgr/database"
	"github.com/AlejandroAM91/athenadb/internal/app/athenadb/datamgr"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// The APIManager manages application API requests.
type APIManager struct {
	httpsrv *http.Server
}

// CreateAPIManager creates and initializes aplication API manager.
func CreateAPIManager(endpoint string, datamgr *datamgr.DataManager) *APIManager {
	// Initialize routes
	router := mux.NewRouter()
	database.Init(router, datamgr)

	// Creates handler
	handler := handlers.LoggingHandler(os.Stdout, router)

	// Create http server
	httpsrv := &http.Server{
		Addr:    endpoint,
		Handler: handler,
	}

	return &APIManager{
		httpsrv: httpsrv,
	}
}

// Start starts http server
func (api APIManager) Start() {
	log.Fatal(api.httpsrv.ListenAndServe())
}
