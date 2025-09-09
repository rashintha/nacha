package types

import (
	"strings"
)

type NachaBlockFiller struct {
	Reserved string // Char Count: 94 | Value: 9s
}

func (b *NachaBlockFiller) Default() {
	b.Reserved = strings.Repeat("9", 94)
}
