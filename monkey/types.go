package monkey

type HeaderType []string
type RowType []string

type style struct {
	textCol  string
	tableCol string
	isLeftAlign bool
}

type Table struct {
	header HeaderType
	rows   []RowType
	styles style
}
