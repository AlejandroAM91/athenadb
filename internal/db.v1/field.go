package db

import (
	"encoding/binary"
	"io"
)

const (
	FieldTypeString uint16 = iota
	FieldMaxBytes   uint16 = 256
)

type Field interface {
	io.WriterTo
	io.ReaderFrom
}

type String string

func (sf *String) ReadFrom(r io.Reader) (int64, error) {
	b1 := make([]byte, 2)
	n1, err := r.Read(b1)
	if err != nil {
		return 0, err
	}

	len := binary.BigEndian.Uint16(b1)
	b2 := make([]byte, int(len))
	n2, err := r.Read(b2)
	if err != nil {
		return 0, err
	}

	*sf = String(b2)
	return int64(n1 + n2), nil
}

func (sf *String) WriteTo(w io.Writer) (int64, error) {
	len := len(*sf) % int(FieldMaxBytes)
	buf := make([]byte, len+2)
	binary.BigEndian.PutUint16(buf[0:2], uint16(len))
	copy(buf[2:], *sf)
	n, err := w.Write(buf)
	return int64(n), err
}
