package table

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Handler manages test calls
func Handler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Printf("Path /%s/%s called\n", vars["database"], vars["table"])
}
