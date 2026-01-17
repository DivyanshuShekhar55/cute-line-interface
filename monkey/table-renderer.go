package monkey

import (
	"cute-line-interface/utils"
	"errors"
	"strings"
)

func renderFirstLine(widths []int, styles style) {
	
	var finalStr string
	finalStr += "\U0000250C"                              // ┌
	finalStr += strings.Repeat("\U00002500", widths[0]-1) // -
	finalStr += "\U0000252C"                              // ┬
	finalStr += strings.Repeat("\U00002500", widths[1]-1) // -
	finalStr += "\U0000252C"                              // ┬
	finalStr += strings.Repeat("\U00002500", widths[2]-1) // -
	finalStr += "\U00002510"
	//  ┐
	
	utils.TurnText(finalStr, styles.tableCol, true, false)

}

func renderMiddleRows(t *Table) {}

func renderLastLine(t *Table) {}

func (t *Table) Render() {
	if t.header == nil {
		err := errors.New("Header missing")
		utils.LogError(err)
		return
	}

	//widths := countTotalColChars(t)

	// print first line

}
