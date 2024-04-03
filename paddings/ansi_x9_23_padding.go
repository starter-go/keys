package paddings

import (
	"fmt"

	"github.com/starter-go/keys"
)

type x923 struct {
}

func (inst *x923) _impl() keys.Padding {
	return inst
}

func (inst *x923) Mode() keys.PaddingMode {
	return keys.PaddingANSIX923
}

func (inst *x923) AddPadding(src []byte, blockSize int) ([]byte, error) {
	return nil, fmt.Errorf("no impl")

}

// RemovePadding ...
func (inst *x923) RemovePadding(src []byte, blockSize int) ([]byte, error) {
	return nil, fmt.Errorf("no impl")
}
