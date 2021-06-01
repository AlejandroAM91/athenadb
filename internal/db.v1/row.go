package db

import (
	"io"
)

type Row struct {
	cols   []Col
	fields []Field
}

func (row *Row) ReadFrom(r io.Reader) (n int64, err error) {
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

func (row *Row) WriteTo(w io.Writer) (n int64, err error) {
	// b1 := make([]byte, 2)
	// binary.BigEndian.PutUint16(b1, uint16(len(*row)))

	// buf := bytes.NewBuffer(b1)
	// for _, row := range *row {
	// 	row.WriteTo(buf)
	// }
	// return buf.WriteTo(w)
	return 0, nil
}
