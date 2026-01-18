package monkey

import "unicode/utf8"

func countTotalColChars(t *Table) []int {

	// no! it's a human comment. i hate AI slop...
	// returns length of longest element in each col, so we can make table cols that >= that col's size
	// assume i have array as:
	// [["name", "age", "email_address"], ["ippo", "23", "ip@kun"], ["fushiguro", "25", "fu@sh"]]
	// colSizes = [3]int, slices would be rows, items are string values
	// we choose max length value as the size decider for each column
	// here we would choose ["fushiguro", "23", "email_address"]

	// UPDATED BUG FIX :
	// removed the len() method cause was having weirdish tables
	// reason was unicode -, |, ... have diff width then alphabets
	// using length counting with runes

	array := convertTableToArray(*t)
	colSizes := make([]int, len(array[0]))

	for _, slice := range array {
		for col, item := range slice {
			width := utf8.RuneCountInString(item) // visual chars
			if width > colSizes[col] {
				colSizes[col] = width
			}
		}
	}
	// add padding for spaces inside each cell: " "+text+" "
	for i := range colSizes {
		colSizes[i] += 2 // 1 left space + 1 right space
	}

	return colSizes
}
