package monkey

type HeaderType []string
type RowType []string

type Style struct {
	textCol  string
	tableCol string
}

type Table struct {
	header HeaderType
	rows   []RowType
	Style  Style
}

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

