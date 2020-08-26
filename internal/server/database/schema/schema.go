package schema

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Handler manages schema calls
func Handler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Printf("Called /%s/schema (database schema)\n", vars["database"])
}
