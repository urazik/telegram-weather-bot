package format

import "strconv"

func FTS0(i float64) string {
	return strconv.FormatFloat(i, 'f', 0, 64)
}

func FTS6(i float64) string {
	return strconv.FormatFloat(i, 'f', 6, 64)
}
