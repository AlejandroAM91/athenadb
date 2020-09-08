package main

import (
	"fmt"

	"github.com/AlejandroAM91/athenadb/internal/app/athenadb/core/model"
	"github.com/AlejandroAM91/athenadb/internal/app/athenadb/core/services"
)

func main() {
	dbsrv := services.CreateDatabaseService()

	database := model.CreateDatabase()
	dbsrv.CreateDatabase("mybookshelf", database)
	fmt.Println("Create database")
	fmt.Println("Service: ", dbsrv)
	fmt.Println("Database: ", database)

	
}
