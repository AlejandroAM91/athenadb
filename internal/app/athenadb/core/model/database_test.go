package model

import (
	"reflect"
	"testing"
)

func TestCreateDatabase(t *testing.T) {
	db := Database{
		colList: []string{},
	}
	res := CreateDatabase([]string{})
	if !reflect.DeepEqual(db, *res) {
		t.Error("Database values mismatch")
	}
}
