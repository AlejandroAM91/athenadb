package main

import (
	"github.com/AlejandroAM91/athenadb/internal/app/athenadb/core/model"
	"github.com/AlejandroAM91/athenadb/internal/app/athenadb/core/services"
)

func main() {
	database := model.CreateDatabase()
	services.CreateDatabase("mybookshelf", database)
}
