package model

import "testing"

func TestCreateDatabase(t *testing.T) {
	db := Database{}
	res := CreateDatabase()
	if db != *res {
		t.Error("Database values mismatch")
	}
}
