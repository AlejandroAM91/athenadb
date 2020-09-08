package athenadb

import (
	"github.com/AlejandroAM91/athenadb/internal/app/athenadb/api"
	"github.com/AlejandroAM91/athenadb/internal/app/athenadb/core/model"
	"github.com/AlejandroAM91/athenadb/internal/app/athenadb/core/services"
)

// The App contains global instance values of the application.
type App struct {
	api      *api.Adapter
	services *services.Services
}

// CreateApp creates and initializes the application.
func CreateApp(endpoint string) App {
	// Create application services manager
	services := services.CreateServices()

	// Create application api adapter
	api := api.CreateAdapter(services, endpoint)

	// Create mock data
	services.GetDatabaseService().CreateDatabase("mybookshelf", model.CreateDatabase())

	return App{
		api:      api,
		services: services,
	}
}

// Start starts the application
func (app App) Start() {
	app.api.Start()
}
