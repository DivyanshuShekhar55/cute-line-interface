package monkey

import (
	"cute-line-interface/utils"
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"
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

	// left border
	finalStr += utils.TurnText("\u2502", styles.tableCol, false, false)

	for i, w := range widths {
		text := ""
		if i < len(values) {
			text = values[i]
		}

		// visual length, consistent with countTotalColChars
		// removed the len() method cause was having weirdish tables
		// reason was unicode -, |, ... have diff width then alphabets
		textWidth := utf8.RuneCountInString(text)

		// width of column - 2 (for |) - text width
		// gives remaining column width to apply as padding
		padding := w - 2 - textWidth
		if padding < 0 {
			padding = 0 // safety guard
		}

		// 1 space + text + padding spaces + 1 space = exactly w cells
		cell := "" + text + strings.Repeat(" ", padding) + " "

		finalStr += utils.TurnText(cell, styles.textCol, false, false)
		finalStr += utils.TurnText("\u2502", styles.tableCol, false, false)
	}

	fmt.Println(finalStr)
}

func (t *Table) Render(textCol, bgCol string) {
	if t.header == nil {
		err := errors.New("Header missing")
		utils.LogError(err)
		return
	}

	styles := style{
		textCol:  textCol,
		tableCol: bgCol,
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
