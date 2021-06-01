package db

import (
	"errors"
	"io"
)

type Table struct {
	/* Metadata */

	names  map[string]int
	schema []FieldMeta

	/* Data */

	rows [][]FieldData
}

func NewTable(schema []FieldMeta) *Table {
	names := make(map[string]int)
	for i, f := range schema {
		names[f.Name()] = i
	}

	return &Table{
		names:  names,
		schema: schema,
	}
}

func (t *Table) Insert(row map[string]FieldData) error {
	nrow := make([]FieldData, len(t.schema))
	for k, v := range row {
		i, ok := t.names[k]
		if !ok {
			return errors.New("inserting field undefined in schema")
		}
		nrow[i] = v
	}
	t.rows = append(t.rows, nrow)
	return nil
}

func (t *Table) Select() ([][]FieldData, error) {
	return t.rows, nil
}

func (t *Table) ReadFrom(r io.Reader) (n int64, err error) {
	// b1 := make([]byte, 2)
	// n1, err := r.Read(b1)
	// if err != nil {
	// 	return 0, err
	// }
	// len := binary.BigEndian.Uint16(b1)

	// row.fields = make([]Field, len)
	// for i := 0; i < int(len); i++ {
	// 	col := row.cols[i]

	// }

	// b1 := make([]byte, 2)
	// n1, err := r.Read(b1)
	// if err != nil {
	// 	return 0, err
	// }
	return 0, nil
}

func (t *Table) WriteTo(w io.Writer) (n int64, err error) {
	// b1 := make([]byte, 2)
	// binary.BigEndian.PutUint16(b1, uint16(len(*row)))

	// buf := bytes.NewBuffer(b1)
	// for _, row := range *row {
	// 	row.WriteTo(buf)
	// }
	// return buf.WriteTo(w)
	return 0, nil
}
