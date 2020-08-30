package athenadb

import (
	"fmt"

	"github.com/AlejandroAM91/athenadb/internal/app/athenadb/apimgr"
	"github.com/AlejandroAM91/athenadb/internal/app/athenadb/datamgr"
)

// The App contains global instance values of the application.
type App struct {
	apimgr *apimgr.APIManager
}

// CreateApp creates and initializes the application http server.
func CreateApp(endpoint string) App {
	// Create application data manager
	datamgr := datamgr.CreateDataManager()
	fmt.Println(datamgr.DatabaseList)

	// Create application data manager
	apimgr := apimgr.CreateAPIManager(endpoint, datamgr)

	return App{
		apimgr: apimgr,
	}
}

// StartServer starts application http server
func (app App) StartServer() {
	app.apimgr.Start()
}
