package services

import (
	"testing"

	"github.com/AlejandroAM91/athenadb/internal/app/athenadb/core/model"
)

func TestCreateDatabaseService(t *testing.T) {
	res := CreateDatabaseService()

	if len(res.databaseMap) != 0 {
		t.Error("Database map should be empty")
	}
}

func TestDatabaseService_CreateDatabase(t *testing.T) {
	db := model.CreateDatabase()
	dbsrv := CreateDatabaseService()
	error := dbsrv.CreateDatabase("test", db)

	if error != nil {
		t.Error("CreateDatabase should not return an error")
	}

	if _, exists := dbsrv.databaseMap["test"]; !exists {
		t.Error("CreateDatabase should add a correct key in databaseMap")
	}

	if res, _ := dbsrv.databaseMap["test"]; res != db {
		t.Error("CreateDatabase should add a correct value in databaseMap")
	}
}

func TestDatabaseService_CreateDatabase_Exists(t *testing.T) {
	dbsrv := CreateDatabaseService()
	dbsrv.CreateDatabase("test", model.CreateDatabase())
	error := dbsrv.CreateDatabase("test", model.CreateDatabase())

	if error == nil {
		t.Error("CreateDatabase should return an error")
	}
}

func TestDatabaseService_RetrieveDatabase(t *testing.T) {
	db := model.CreateDatabase()
	dbsrv := CreateDatabaseService()
	dbsrv.CreateDatabase("test", db)

	res, error := dbsrv.RetrieveDatabase("test")

	if error != nil {
		t.Error("RetrieveDatabase should not return an error")
	}

	if res == nil {
		t.Error("RetrieveDatabase should return a not nil database value")
	}

	if res != db {
		t.Error("RetrieveDatabase should return a correct database value")
	}
}

func TestDatabaseService_RetrieveDatabase_NotExists(t *testing.T) {
	dbsrv := CreateDatabaseService()
	res, error := dbsrv.RetrieveDatabase("test")

	if error == nil {
		t.Error("RetrieveDatabase should return an error")
	}

	if res != nil {
		t.Error("RetrieveDatabase should return a nil database value")
	}
}

func TestDatabaseService_UpdateDatabase(t *testing.T) {
	db := model.CreateDatabase()
	dbsrv := CreateDatabaseService()
	dbsrv.CreateDatabase("test", model.CreateDatabase())
	error := dbsrv.UpdateDatabase("test", db)

	if error != nil {
		t.Error("UpdateDatabase should not return an error")
	}

	if res, _ := dbsrv.databaseMap["test"]; db != res {
		t.Error("UpdateDatabase should update a correct value in databaseMap")
	}
}

func TestDatabaseService_UpdateDatabase_NotExists(t *testing.T) {
	dbsrv := CreateDatabaseService()
	error := dbsrv.UpdateDatabase("test", model.CreateDatabase())

	if error == nil {
		t.Error("UpdateDatabase should return an error")
	}
}

func TestDatabaseService_DeleteDatabase(t *testing.T) {
	dbsrv := CreateDatabaseService()
	dbsrv.CreateDatabase("test", model.CreateDatabase())
	error := dbsrv.DeleteDatabase("test")

	if error != nil {
		t.Error("DeleteDatabase should not return an error")
	}

	if _, exists := dbsrv.databaseMap["test"]; exists {
		t.Error("DeleteDatabase should remove a correct key in databaseMap")
	}
}

func TestDatabaseService_DeleteDatabase_NotExists(t *testing.T) {
	dbsrv := CreateDatabaseService()
	error := dbsrv.DeleteDatabase("test")

	if error == nil {
		t.Error("DeleteDatabase should return an error")
	}
}
