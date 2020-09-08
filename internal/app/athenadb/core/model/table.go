package model

// The Table contains the information about a table
type Table struct {
	rowList    [][]string
	columnList []string
}

// CreateTable creates and initializes a table.
func CreateTable() *Table {
	return &Table{
		rowList:    make([][]string, 0),
		columnList: make([]string, 0),
	}
}

// AddColumn adds a new column to the table schema.
func (table *Table) AddColumn(column ...string) {
	table.columnList = append(table.columnList, column...)
}

// AddColumnList append a list of columns to the table schema.
func (table *Table) AddColumnList(columnList []string) {
	table.columnList = append(table.columnList, columnList...)
}

// AddRow adds a new row to the table.
func (table *Table) AddRow(row []string) {
	table.rowList = append(table.rowList, row)
}

// GetAll get all table rows.
func (table Table) GetAll() [][]string {
	return table.rowList
}
