package database

import (
	"github.com/AlejandroAM91/athenadb/internal/app/athenadb/apimgr/database/schema"
	"github.com/AlejandroAM91/athenadb/internal/app/athenadb/apimgr/database/table"
	"github.com/AlejandroAM91/athenadb/internal/app/athenadb/datamgr"
	"github.com/gorilla/mux"
)

// Init initialize test routes
func Init(router *mux.Router, datamgr *datamgr.DataManager) {

	// Initialize handlers
	router.Handle("/{database}/schema", schema.CreateSchema(datamgr))
	router.HandleFunc("/{database}/{table}", table.Handler)
	router.HandleFunc("/{database}/{table}/schema", table.SchemaHandler)
}
