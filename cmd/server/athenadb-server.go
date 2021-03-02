package main

import (
	"fmt"
	"net/http"

	"github.com/AlejandroAM91/athenadb/internal/app/server/adapters/http/server"
	"github.com/AlejandroAM91/athenadb/internal/app/server/core/services"
	"github.com/AlejandroAM91/athenadb/pkg/core/model"
	"github.com/go-chi/chi/v5"
)

type dbkvs struct {
	dbmap map[string]model.DB
}

func (s dbkvs) GetAll() ([]model.DB, error) {
	dblist := make([]model.DB, 0)
	for _, v := range s.dbmap {
		dblist = append(dblist, v)
	}
	return dblist, nil
}

func main() {
	fmt.Println("Starting athenadb")

	sa := &dbkvs{dbmap: make(map[string]model.DB)}
	sa.dbmap["books"] = model.DB{Name: "El nombre del viento"}
	dbsrv := services.NewDBSrv(sa)

	r := chi.NewRouter()
	r.Mount("/admin/database", server.NewDBHdl(dbsrv))
	http.ListenAndServe(":8080", r)
}
