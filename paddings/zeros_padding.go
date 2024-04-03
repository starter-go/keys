package paddings

import (
	"fmt"

	"github.com/starter-go/keys"
)

type zeros struct {
}

func (inst *zeros) _impl() keys.Padding {
	return inst
}

func (inst *zeros) Mode() keys.PaddingMode {
	return keys.PaddingZeros
}

func (inst *zeros) AddPadding(src []byte, blockSize int) ([]byte, error) {
	if blockSize < 1 {
		return nil, fmt.Errorf("ZeroPadding: bad blockSize: %d", blockSize)
	}
	dst := src
	pLen := computePaddingLength(src, blockSize)
	for count := 0; count < pLen; count++ {
		dst = append(dst, 0)
	}
	return dst, nil
}

func (inst *zeros) RemovePadding(src []byte, blockSize int) ([]byte, error) {
	srcLen := len(src)
	start := srcLen - blockSize
	end := srcLen - 1
	if start < 0 {
		start = 0
	}
	i := end
	for ; i > start; i-- {
		b := src[i]
		if b != 0 {
			break
		}
	}
	return src[0 : i+1], nil
}
