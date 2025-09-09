package util

import (
	"fmt"
	"strings"
)

// ToFixedWidthString returns a string with the specified width.
// If the string is longer than the width, it will be truncated.
// If alignRight is true, the string will be right aligned.
// Otherwise, it will be left aligned.
// Blank spaces will be added to the string if it is shorter than the width.
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

// ToFixedWidthZeroString returns a string with the specified width.
// If the string is longer than the width, it will be truncated.
// Zeros will be added to the string if it is shorter than the width.
func ToFixedWidthZeroString(s string, width int) string {
	if len(s) > width {
		s = s[:width] // truncate
	}

	return strings.Repeat("0", width-len(s)) + s
}
