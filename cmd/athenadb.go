package main

import "github.com/AlejandroAM91/athenadb/internal/app/athenadb"

func main() {
	app := athenadb.CreateApp(":80")
	app.Start()
}
