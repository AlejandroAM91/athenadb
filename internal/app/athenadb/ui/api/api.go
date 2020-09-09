package api

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/AlejandroAM91/athenadb/internal/app/athenadb/ui/api/admin"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// The API manages application API requests.
type API struct {
    handler  http.Handler
    server   *http.Server
}

// NewAPI creates and initializes aplication API.
func NewAPI() *API {
    // Initialize handler
    router := initRoutes()
    handler := initHandler(router)

    return &API{
        handler: handler,
    }
}

// Start starts http server
func (api *API) Start(endpoint string) {
    // Stops previous server
    api.Stop()

    // Create http server
    api.server = &http.Server{
        Addr:    endpoint,
        Handler: api.handler,
    }
    go func() {
        if err := api.server.ListenAndServe(); err != http.ErrServerClosed {
            log.Fatal(err)
        }
    }()
}

// Stop stops http server
func (api API) Stop() {
    // Stops previous server
    if api.server != nil {
        api.server.Shutdown(context.Background())
    }
}

func initHandler(router *mux.Router) http.Handler {
	return handlers.LoggingHandler(os.Stdout, router)
}

func initRoutes() *mux.Router {
	router := mux.NewRouter()
	router.Handle("/admin/database", admin.CreateDatabaseController())
	return router
}
