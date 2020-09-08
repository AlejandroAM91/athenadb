package api

import (
	"log"
	"net/http"
	"os"

	"github.com/AlejandroAM91/athenadb/internal/app/athenadb/core/services"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// The Adapter manages application API requests.
type Adapter struct {
	server   *http.Server
	services *services.Services
}

// CreateAdapter creates and initializes aplication API adapter.
func CreateAdapter(services *services.Services, endpoint string) *Adapter {
	// Initialize routes
	router := mux.NewRouter()
	router.Handle("/{database}", CreateDatabaseController(services))

	// Creates handlers
	handler := handlers.LoggingHandler(os.Stdout, router)

	// Create http server
	server := &http.Server{
		Addr:    endpoint,
		Handler: handler,
	}

	return &Adapter{
		server:   server,
		services: services,
	}
}

// Start starts http server
func (adapter Adapter) Start() {
	log.Fatal(adapter.server.ListenAndServe())
}
