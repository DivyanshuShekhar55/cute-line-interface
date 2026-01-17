package monkey

func NewTable() *Table {
	return &Table{}
}

func (t *Table) Header(vals []string) *Table {
	t.header = HeaderType(vals)
	return t
}

func (t *Table) Row(vals []string) *Table {
	t.rows = append(t.rows, vals)
	return t
}

func (t *Table) Style(textCol, tableCol string, isLeftAlign bool) *Table {
	t.styles.tableCol = tableCol
	t.styles.textCol = textCol
	t.styles = style{tableCol, tableCol, isLeftAlign}
	return t
}

func convertTableToArray(t Table) [][]string {
	// we convert all values to strings
	// rows+1 for rows+header
	array := make([][]string, 0, len(t.rows)+1)

	// add header
	array = append(array, []string(t.header))

	// add rows
	for _, row := range t.rows {
		array = append(array, []string(row))
	}

	return array
}
