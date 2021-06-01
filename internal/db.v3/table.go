package db

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
)

const (
	fieldKindString uint16 = iota
)

type Table struct {
	rows [][]interface{}
}

func NewTable() *Table {
	return &Table{}
}

func (t *Table) Insert(row []interface{}) error {
	t.rows = append(t.rows, row)
	return nil
}

func (t *Table) Select() ([][]interface{}, error) {
	return t.rows, nil
}

func (t *Table) ReadFrom(r io.Reader) (int64, error) {
	b1 := make([]byte, 2)
	n, err := r.Read(b1)
	if err != nil {
		return 0, err
	}

	len := int(binary.BigEndian.Uint16(b1))
	t.rows = make([][]interface{}, len)
	for i := 0; i < len; i++ {
		row, n2, err := readRow(r)
		if err != nil {
			return 0, err
		}
		t.rows = append(t.rows, row)
		n += n2
	}
	return int64(n), nil
}

func (t *Table) WriteTo(w io.Writer) (int64, error) {
	b1 := make([]byte, 2)
	binary.BigEndian.PutUint16(b1, uint16(len(t.rows)))

	buf := bytes.NewBuffer(b1)
	for _, row := range t.rows {
		writeRow(buf, row)
	}
	return buf.WriteTo(w)
}

func readRow(r io.Reader) ([]interface{}, int, error) {
	// Reads row len
	// b := make([]byte, 2)
	// n, err := r.Read(b)
	// if err != nil {
	// 	return nil, 0, err
	// }
	// len := int(binary.BigEndian.Uint16(b))

	// Reads each field
	// for i := 0; i < len; i++ {
	// 	n2, err := r.Read(b)
	// 	if err != nil {
	// 		return nil, 0, err
	// 	}
	// 	len := int(binary.BigEndian.Uint16(b))

	// 	switch v := field.(type) {
	// 	case string:
	// 		writeString(w, v)
	// 	default:
	// 		return errors.New("writing unsupported type")
	// 	}
	// }
	return nil, 0, nil
}

func readField(r io.Reader) (data interface{}, n int, err error) {
	// switch v := field.(type) {
	// case string:
	// 	n, err = writeString(w, v)
	// default:
	// 	return 0, errors.New("writing unsupported type")
	// }
	// return n, err
	return nil, 0, nil
}

func writeRow(w io.Writer, row []interface{}) (n int, err error) {
	// Writes row len
	if n, err = writeValue(w, uint16(len(row))); err != nil {
		return 0, err
	}

	// Writes each field
	for _, field := range row {
		var n1 int
		if n1, err = writeField(w, field); err != nil {
			return 0, err
		}
		n += n1
	}
	return n, nil
}

func writeField(w io.Writer, value interface{}) (n int, err error) {
	switch value.(type) {
	case string:
		n, err = writeValue(w, fieldKindString)
	default:
		return 0, errors.New("writing unsupported type")
	}

	var n2 int
	if err == nil {
		n2, err = writeValue(w, value)
	}
	return n + n2, err
}

func writeValue(w io.Writer, value interface{}) (int, error) {
	var b []byte
	switch v := value.(type) {
	case uint16:
		b = make([]byte, 2)
		binary.BigEndian.PutUint16(b, v)
	case string:
		b = make([]byte, 2+len(v))
		binary.BigEndian.PutUint16(b[:2], uint16(len(v)))
		copy(b[2:], v)
	default:
		return 0, errors.New("writing unsupported type")
	}
	return w.Write(b)
}
