package main

import (
	"github.com/AlejandroAM91/athenadb/internal/app/athenadb/api"
	"github.com/AlejandroAM91/athenadb/internal/pkg/dimgr"
)

func main() {
	dimgr := dimgr.GetManager()
	api := dimgr.GetDependency("api").(*api.Adapter)
	api.SyncStart(":80")
}
