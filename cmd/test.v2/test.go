package main

import (
	"fmt"

	"github.com/AlejandroAM91/athenadb/internal/db.v2"
)

func main() {
	// f, err := os.CreateTemp(".tmp", "*.adb")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()

	schema := []db.FieldMeta{
		db.CreateStringField("Name", 50),
		db.CreateStringField("Author", 50),
	}
	t := db.NewTable(schema)

	data := []map[string]db.FieldData{
		{
			"Name":   db.String("El dragon de su majestad"),
			"Author": db.String("Naomi Novik"),
		},
		{
			"Name":   db.String("El trono de jade"),
			"Author": db.String("Naomi Novik"),
		},
	}

	for _, v := range data {
		t.Insert(v)
	}
	fmt.Println(*t)
	// t.WriteTo(f)
	// t.Create(db.Row{[]string{"Hola"}})

}
