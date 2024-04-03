package paddings

import (
	"fmt"

	"github.com/starter-go/keys"
)

type iso10126 struct {
}

func (inst *iso10126) _impl() keys.Padding {
	return inst
}

func (inst *iso10126) Mode() keys.PaddingMode {
	return keys.PaddingISO10126
}

func (inst *iso10126) AddPadding(src []byte, blockSize int) ([]byte, error) {
	return nil, fmt.Errorf("no impl")

}

// RemovePadding ...
func (inst *iso10126) RemovePadding(src []byte, blockSize int) ([]byte, error) {
	return nil, fmt.Errorf("no impl")
}
