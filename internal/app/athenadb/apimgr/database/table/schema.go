package table

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// SchemaHandler manages table calls
func SchemaHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Printf("Called /%s/%s/schema (table schema)\n", vars["database"], vars["table"])
}
