package db

import (
	"encoding/binary"
	"io"
)

type Col struct {
	kind uint16
	size uint16
}

func (c Col) WriteTo(w io.Writer) (int64, error) {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint16(buf[:2], c.kind)
	binary.BigEndian.PutUint16(buf[2:], c.size)
	n, err := w.Write(buf)
	return int64(n), err
}

func ColString(size int) Col {
	return Col{
		kind: FieldTypeString,
		size: uint16(size),
	}
}
