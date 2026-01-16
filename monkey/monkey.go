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

func (t *Table) Style(textCol, tableCol string) *Table {
	t.styles.tableCol = tableCol
	t.styles.textCol = textCol
	return t
}
