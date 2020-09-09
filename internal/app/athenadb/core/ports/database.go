package ports

import (
	"github.com/AlejandroAM91/athenadb/internal/app/athenadb/core/internal/services"
	"github.com/AlejandroAM91/athenadb/internal/app/athenadb/core/model"
	"github.com/AlejandroAM91/athenadb/internal/pkg/dimgr"
)

// CreateDatabasePort port to create a database.
type CreateDatabasePort interface {
	CreateDatabase(name string, database *model.Database) error
}

func init() {
	dimgr := dimgr.GetManager()
	dbsrv := services.CreateDatabaseService()
	dimgr.AddDependency("CreateDatabasePort", CreateDatabasePort(dbsrv))
}
