package paddings

import "github.com/starter-go/keys"

////////////////////////////////////////////////////////////////////////////////

// NewNoPadding ...
func NewNoPadding() keys.Padding {
	return new(nop)
}

// NewPKCS5Padding ...
func NewPKCS5Padding() keys.Padding {
	return new(pkcs5)
}

// NewPKCS7Padding ...
func NewPKCS7Padding() keys.Padding {
	return new(pkcs7)
}

// NewZerosPadding ...
func NewZerosPadding() keys.Padding {
	return new(zeros)
}

// NewANSIX923Padding ...
func NewANSIX923Padding() keys.Padding {
	return new(x923)
}

// NewISO10126Padding ...
func NewISO10126Padding() keys.Padding {
	return new(iso10126)
}

////////////////////////////////////////////////////////////////////////////////

// NewPadding 根据传入的模式创建对应的填充器
func NewPadding(mode keys.PaddingMode) keys.Padding {
	switch mode {
	case keys.PaddingPKCS7:
		return NewPKCS7Padding()

	case keys.PaddingPKCS5:
		return NewPKCS5Padding()

	case keys.PaddingANSIX923:
		return NewANSIX923Padding()

	case keys.PaddingISO10126:
		return NewISO10126Padding()

	case keys.PaddingZeros:
		return NewZerosPadding()

	case keys.NoPadding:
		return NewNoPadding()
	}
	// 默认以 PKCS7 填充
	return NewPKCS7Padding()
}

////////////////////////////////////////////////////////////////////////////////

// 计算默认的填充长度
func computePaddingLength(src []byte, blockSize int) int {
	if blockSize < 1 {
		return 0
	}
	rawSize := len(src)
	return blockSize - (rawSize % blockSize)
}
