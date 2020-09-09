package api

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/AlejandroAM91/athenadb/internal/pkg/dimgr"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// The Adapter manages application API requests.
type Adapter struct {
	serverClosed chan struct{}
	handler      http.Handler
	server       *http.Server
}

// CreateAdapter creates and initializes aplication API adapter.
func CreateAdapter() *Adapter {
	// Initialize routes
	router := mux.NewRouter()
	initRoutes(router)

	// Creates handlers
	handler := handlers.LoggingHandler(os.Stdout, router)

	return &Adapter{
		serverClosed: make(chan struct{}),
		handler:      handler,
	}
}

// Start starts http server
func (adapter *Adapter) Start(endpoint string) {
	// Stops previous server
	adapter.Stop()

	// Create http server
	adapter.server = &http.Server{
		Addr:    endpoint,
		Handler: adapter.handler,
	}
	go func() {
		if err := adapter.server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatal(err)
		}
		close(adapter.serverClosed)
	}()
}

// Stop stops http server
func (adapter Adapter) Stop() {
	// Stops previous server
	if adapter.server != nil {
		adapter.server.Shutdown(context.Background())
	}
}

// Sync wait until server is stopped
func (adapter Adapter) Sync() {
	<-adapter.serverClosed
}

// SyncStart starts http server syncronous
func (adapter Adapter) SyncStart(endpoint string) {
	adapter.Start(endpoint)
	adapter.Sync()
}

func init() {
	dimgr := dimgr.GetManager()
	dimgr.AddDependency("api", CreateAdapter())
}

func initRoutes(router *mux.Router) {
	router.Handle("/{database}", CreateDatabaseController())
}
