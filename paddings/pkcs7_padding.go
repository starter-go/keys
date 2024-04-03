package paddings

import (
	"fmt"

	"github.com/starter-go/keys"
)

type pkcs7 struct {
}

func (inst *pkcs7) _impl() keys.Padding {
	return inst
}

func (inst *pkcs7) Mode() keys.PaddingMode {
	return keys.PaddingPKCS7
}

func (inst *pkcs7) AddPadding(src []byte, blockSize int) ([]byte, error) {
	dst := src
	pSize := computePaddingLength(src, blockSize)
	if 0 < pSize && pSize < 256 {
		b := byte(pSize)
		for count := 0; count < pSize; count++ {
			dst = append(dst, b)
		}
	} else {
		return nil, fmt.Errorf("pkcs7padding: bad padding size: %d", pSize)
	}
	return dst, nil
}

func (inst *pkcs7) RemovePadding(src []byte, blockSize int) ([]byte, error) {

	// check src length

	srcLen := len(src)

	if srcLen < 1 {
		return nil, fmt.Errorf("pkcs7padding: bad src size [%d] for depad", srcLen)
	}

	if blockSize < 1 {
		return nil, fmt.Errorf("pkcs7padding: bad block size [%d] for depad", blockSize)
	}

	if (srcLen % blockSize) != 0 {
		f := "pkcs7padding: bad size for depad, [src_data_size:%d block_size:%d]"
		return nil, fmt.Errorf(f, srcLen, blockSize)
	}

	// do depad

	pSize := src[srcLen-1]
	dstLen := srcLen - int(pSize)
	dst := src[0:dstLen]
	return dst, nil
}
