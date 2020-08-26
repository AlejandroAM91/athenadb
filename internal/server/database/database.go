package database

import (
	"github.com/AlejandroAM91/athenadb/internal/server/database/schema"
	"github.com/AlejandroAM91/athenadb/internal/server/database/table"
	"github.com/gorilla/mux"
)

// Init initialize test routes
func Init(router *mux.Router) {

	// Initialize handlers
	router.HandleFunc("/{database}/schema", schema.Handler)
	router.HandleFunc("/{database}/{table}", table.Handler)
	router.HandleFunc("/{database}/{table}/schema", table.SchemaHandler)
}
