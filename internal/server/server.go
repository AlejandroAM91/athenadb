package server

import (
	"net/http"

	"github.com/AlejandroAM91/athenadb/internal/server/test"
	"github.com/gorilla/mux"
)

// Start initializes and starts application server
func Start(endpoint string) {
	// Initialize routes
	router := mux.NewRouter()
	test.Init(router)

	// Start server
	http.ListenAndServe(endpoint, router)
}
