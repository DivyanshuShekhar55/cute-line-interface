package monkey

type HeaderType []string
type RowType []string

type style struct {
	textCol  string
	tableCol string
}

type Table struct {
	header HeaderType
	rows   []RowType
	styles style
}