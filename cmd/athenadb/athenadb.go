package main

import (
	"fmt"

	"github.com/AlejandroAM91/athenadb/internal/app/athenadb/services"
	"github.com/AlejandroAM91/athenadb/pkg/domain/model"
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
	sa := &dbkvs{
		dbmap: make(map[string]model.DB),
	}
	sa.dbmap["books"] = model.DB{Name: "El nombre del viento"}
	dbsrv := services.NewDBSrv(sa)
	dblist, _ := dbsrv.GetAll()
	for _, db := range dblist {
		fmt.Println(db)
	}
}
