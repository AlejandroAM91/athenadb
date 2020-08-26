package table

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Handler manages table calls
func Handler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Printf("Called /%s/%s (table data)\n", vars["database"], vars["table"])
}
