package util

import "fmt"

func ToFixedWidthString(s string, width int, alignRight bool) string {
	if len(s) > width {
		s = s[:width] // truncate
	}

	if alignRight {
		return fmt.Sprintf("%*s", width, s)
	} else {
		return fmt.Sprintf("%-*s", width, s)
	}
}
