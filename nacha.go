package nacha

import (
	"github.com/rashintha/nacha/types"
)

func NewFile() *types.NachaFile {
	file := &types.NachaFile{}
	file.Header.Default()
	file.Control.Default()
	return file
}
