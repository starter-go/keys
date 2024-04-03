package paddings

import (
	"github.com/starter-go/keys"
)

// pkcs5 是 blockSize 固定为 8 的 pkcs7
type pkcs5 struct {
	p7 pkcs7
}

func (inst *pkcs5) _impl() keys.Padding {
	return inst
}

func (inst *pkcs5) Mode() keys.PaddingMode {
	return keys.PaddingPKCS5
}

func (inst *pkcs5) AddPadding(src []byte, blockSize int) ([]byte, error) {
	blockSize = 8
	return inst.p7.AddPadding(src, blockSize)
}

func (inst *pkcs5) RemovePadding(src []byte, blockSize int) ([]byte, error) {
	blockSize = 8
	return inst.p7.RemovePadding(src, blockSize)
}
