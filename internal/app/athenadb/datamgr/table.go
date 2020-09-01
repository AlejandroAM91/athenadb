package datamgr

// The Table contains the information about a table
type Table struct {
	rows [][][]byte
}

// CreateTable creates and initializes a table.
func CreateTable() *Table {
	return &Table{rows: make([][][]byte, 0)}
}

// AddRow adds a new row to the table.
func (table *Table) AddRow(row [][]byte) {
	table.rows = append(table.rows, row)
}

// GetAll get all table rows.
func (table Table) GetAll() [][][]byte {
	return table.rows
}
