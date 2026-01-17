package monkey

func countTotalColChars(t *Table) []int {

	// no! it's a human comment. i hate AI slop...
	// returns length of longest element in each col, so we can make table cols that >= that col's size 
	// assume i have array as:
	// [["name", "age", "email_address"], ["ippo", "23", "ip@kun"], ["fushiguro", "25", "fu@sh"]]
	// colSizes = [3]int, slices would be rows, items are string values
	// we choose max length value as the size decider for each column
	// here we would choose ["fushiguro", "23", "email_address"]

	array := convertTableToArray(*t)
	colSizes := make([]int, len(array[0]))

	for _, slice := range array {
		for col, item := range slice {
			if len(item) > colSizes[col] {
				colSizes[col] = len(item)
			}
		}
	}
	return colSizes
}
