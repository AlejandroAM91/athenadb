package server

import (
	"log"
	"net/http"

	"github.com/AlejandroAM91/athenadb/internal/server/database"
	"github.com/gorilla/mux"
)

// Start initializes and starts application server
func Start(endpoint string) {
	// Initialize routes
	router := mux.NewRouter()
	database.Init(router)

	// Start server
	server := &http.Server{
		Addr:    endpoint,
		Handler: router,
	}
	log.Fatal(server.ListenAndServe())
}
