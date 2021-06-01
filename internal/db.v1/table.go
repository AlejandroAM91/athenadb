package db

import (
	"bytes"
	"encoding/binary"
	"io"
)

type Table struct {
	cols []Col
	rows []Row
}

func NewTable(cols []Col) *Table {
	return &Table{
		cols: cols,
		rows: make([]Row, 0),
	}
}

func (t *Table) WriteTo(w io.Writer) (int64, error) {
	hdr := make([]byte, 4)
	binary.BigEndian.PutUint16(hdr[:2], uint16(len(t.cols)))
	// TODO: copy rows size

	buf := bytes.NewBuffer(hdr)
	for _, col := range t.cols {
		col.WriteTo(buf)
	}

	// for _, row := range t.rows {
	// row.WriteTo(buf)
	// }

	return buf.WriteTo(w)
}

// func (t *Table) Create(row Row) (err error) {
// 	t.rw.Seek(0, io.SeekEnd)
// 	_, err = row.WriteTo(t.rw)
// 	return err
// }

// func (t *Table) Retrieve() (rows []Row, err error) {

// 	return nil, nil
// }
