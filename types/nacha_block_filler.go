package types

import (
	"strings"
)

// NachaBlockFiller represents the NACHA Block Filler (Type 9)
type NachaBlockFiller struct {
	Reserved string // Char Count: 94 | Value: 9s
}

// Default sets the default values for the NachaBlockFiller
func (b *NachaBlockFiller) Default() {
	b.Reserved = strings.Repeat("9", 94)
}
