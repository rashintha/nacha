package nacha

import (
	"github.com/rashintha/nacha/types"
)

// NewFile creates a new NACHA file
func NewFile() *types.NachaFile {
	file := &types.NachaFile{}
	file.Header.Default()
	file.Control.Default()
	return file
}
