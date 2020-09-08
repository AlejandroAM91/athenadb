package main

import (
	"fmt"

	"github.com/AlejandroAM91/athenadb/internal/app/athenadb/core/model"
)

func main() {
	// app := athenadb.CreateApp(":80")
	// app.StartServer()
	cols := []model.Col{
		model.Col{Name: "_id"},
	}
	table := model.CreateTable(cols)
	database := model.CreateDatabase()
	database.AddTable("book", table)

	table, _ = database.GetTable("book")
	fmt.Println(table.GetAllColumns())
}
