package db

const (
	fieldKindString uint16 = iota
)

type FieldData interface{}
type FieldMeta interface {
	Name() string
}

type fieldMeta struct {
	kind uint16
	name string
	size uint16
}

func CreateStringField(name string, size int) FieldMeta {
	return &fieldMeta{
		kind: fieldKindString,
		name: name,
		size: uint16(size),
	}
}

func (f *fieldMeta) Name() string {
	return f.name
}

type String string
