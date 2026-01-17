package monkey

import (
	"cute-line-interface/utils"
	"errors"
	"fmt"
	"strings"
)

func renderFirstLine(widths []int, styles style) {
	var finalStr string

	finalStr += "\u250C" // ┌

	for i, w := range widths {
		finalStr += strings.Repeat("\u2500", w-1) // ─

		if i < len(widths)-1 {
			finalStr += "\u252C" // ┬
		}
	}

	finalStr += "\u2510" // ┐

	fmt.Println(utils.TurnText(finalStr, styles.tableCol, true, false))
}

func renderMiddleLine(widths []int, styles style) {
	var finalStr string

	finalStr += "\u251C" // ├

	for i, w := range widths {
		finalStr += strings.Repeat("\u2500", w-1) // ─

		if i < len(widths)-1 {
			finalStr += "\u253C" // ┼
		}
	}

	finalStr += "\u2524" // ┤

	fmt.Println(utils.TurnText(finalStr, styles.tableCol, true, false))
}

func renderLastLine(widths []int, styles style) {
	var finalStr string

	finalStr += "\u2514" // └

	for i, w := range widths {
		finalStr += strings.Repeat("\u2500", w-1) // ─

		if i < len(widths)-1 {
			finalStr += "\u2534" // ┴
		}
	}

	finalStr += "\u2518" // ┘

	fmt.Println(utils.TurnText(finalStr, styles.tableCol, true, false))
}

func (t *Table) Render() {
	if t.header == nil {
		err := errors.New("Header missing")
		utils.LogError(err)
		return
	}

	//widths := countTotalColChars(t)

	// print first line

}
