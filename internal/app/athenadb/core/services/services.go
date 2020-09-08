package services

// The Services contains a list of services can be injected
type Services struct {
	dbsrv *DatabaseService
}

// CreateServices creates and initializes the application services.
func CreateServices() *Services {
	return &Services{
		dbsrv: CreateDatabaseService(),
	}
}

// GetDatabaseService returns database service instance.
func (srvmgr Services) GetDatabaseService() *DatabaseService {
	return srvmgr.dbsrv
}
