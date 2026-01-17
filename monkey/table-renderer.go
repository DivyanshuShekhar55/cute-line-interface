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

func renderTextLine(values []string, widths []int, styles style) {
	var finalStr string

	finalStr += "\u2502" // │ left boundary

	for i, w := range widths {
		text := ""
		if i < len(values) {
			text = values[i]
		}

		// trim if overflow
		if len(text) > w-2 {
			text = text[:w-2]
		}

		// pad text
		padding := w - 2 - len(text)
		cell := " " + text + strings.Repeat(" ", padding) + " "

		finalStr += utils.TurnText(cell, styles.textCol, false, false)

		finalStr += "\u2502" // │ column boundary
	}

	fmt.Println(finalStr)
}

func (t *Table) Render(styles style) {
	if t.header == nil {
		err := errors.New("Header missing")
		utils.LogError(err)
		return
	}

	widths := countTotalColChars(t)

	// top border
	renderFirstLine(widths, styles)

	// header row
	renderTextLine(t.header, widths, styles)

	// header separator
	renderMiddleLine(widths, styles)

	// body rows
	for _, row := range t.rows {
		renderTextLine(row, widths, styles)
	}

	// bottom border
	renderLastLine(widths, styles)
}
