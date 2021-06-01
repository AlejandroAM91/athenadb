package main

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/AlejandroAM91/athenadb/internal/db"
)

func main() {
	buf := new(bytes.Buffer)
	mw := io.MultiWriter(os.Stdout, buf)
	t1 := db.NewTable()

	data := [][]interface{}{
		{"El dragon de su majestad", "Naomi Novik"},
		{"El trono de jade", "Naomi Novik"},
	}

	for _, v := range data {
		t1.Insert(v)
	}
	fmt.Println("======= Table 1 =======")
	fmt.Println(*t1)
	fmt.Println()

	fmt.Println("======== Bytes ========")
	t1.WriteTo(mw)
	fmt.Println()

	t2 := db.NewTable()
	t2.ReadFrom(buf)
	fmt.Println("======= Table 2 =======")
	fmt.Println(*t2)

	// t.Create(db.Row{[]string{"Hola"}})

}
