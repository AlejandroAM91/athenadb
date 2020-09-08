package services

// The ServiceManager contains a list of services can be injected
type ServiceManager struct {
	dbsrv *DatabaseService
}

// CreateServiceManager creates and initializes the service manager.
func CreateServiceManager() *ServiceManager {
	return &ServiceManager{
		dbsrv: CreateDatabaseService(),
	}
}

// GetDatabaseService returns database service instance.
func (srvmgr ServiceManager) GetDatabaseService() *DatabaseService {
	return srvmgr.dbsrv
}
