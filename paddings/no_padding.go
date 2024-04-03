package paddings

import (
	"fmt"

	"github.com/starter-go/keys"
)

// 这个模式不填充任何内容，只检查数据块长度是否符合要求
type nop struct {
}

func (inst *nop) _impl() keys.Padding {
	return inst
}

func (inst *nop) Mode() keys.PaddingMode {
	return keys.NoPadding
}

func (inst *nop) check(src []byte, blockSize int) ([]byte, error) {
	if blockSize < 1 {
		return nil, fmt.Errorf("NoPadding: bad block_size: %d", blockSize)
	}
	size := len(src)
	if (size % blockSize) == 0 {
		return src, nil
	}
	f := "NoPadding: bad data length [data_length:%d block_size:%d]"
	return nil, fmt.Errorf(f, size, blockSize)
}

// Pad ...
func (inst *nop) AddPadding(src []byte, blockSize int) ([]byte, error) {
	return inst.check(src, blockSize)
}

// RemovePadding ...
func (inst *nop) RemovePadding(src []byte, blockSize int) ([]byte, error) {
	return inst.check(src, blockSize)
}
