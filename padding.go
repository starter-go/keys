package keys

// PaddingMode 代表一种填充模式
type PaddingMode int

// 定义一些常用的填充模式
const (
	PaddingModeNone PaddingMode = iota

	// for secret-key
	NoPadding
	PaddingPKCS7
	PaddingPKCS5
	PaddingZeros
	PaddingISO10126
	PaddingANSIX923

	// for RSA
	PaddingOAEP
	PaddingPKCS1v15
	PaddingSessionKey
	PaddingPSS
)

func (p PaddingMode) String() string {
	switch p {
	case NoPadding:
		return "NoPadding"
	case PaddingPKCS5:
		return "PKCS5Padding"
	case PaddingPKCS7:
		return "PKCS7Padding"
	case PaddingZeros:
		return "ZerosPadding"
	case PaddingISO10126:
		return "ISO10126Padding"
	case PaddingANSIX923:
		return "ANSIX923Padding"
	}
	return "PaddingModeUndefined"
}

// Padding 是填充算法的接口
type Padding interface {
	Mode() PaddingMode

	AddPadding(src []byte, blockSize int) ([]byte, error)

	RemovePadding(src []byte, blockSize int) ([]byte, error)
}
