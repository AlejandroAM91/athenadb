package model

// The Table contains the information about a table
type Table struct {
	columnList []Col
}

// CreateTable creates and initializes a table.
func CreateTable(columnList []Col) *Table {
	return &Table{
		columnList: columnList,
	}
}

// GetAllColumns retreive all table column descriptions.
func (table Table) GetAllColumns() []Col {
	return table.columnList
}
