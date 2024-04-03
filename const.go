package keys

////////////////////////////////////////////////////////////////////////////////

// CipherMethod 表示密钥的操作方法
type CipherMethod int

// 定义密钥的操作方法
const (
	CipherMethodNone CipherMethod = iota

	CipherMethodDecrypt
	CipherMethodEncrypt
)

func (m CipherMethod) String() string {
	switch m {
	case CipherMethodDecrypt:
		return "Decrypt"
	case CipherMethodEncrypt:
		return "Encrypt"
	}
	return "CipherMethodUndefined"
}

////////////////////////////////////////////////////////////////////////////////

// FlowMode 表示密钥的串联模式
type FlowMode int

// 定义密钥的串联模式
const (
	FlowModeNone FlowMode = iota

	// Cipher Block Chaining
	FlowModeCBC

	// Cipher Feed Back
	FlowModeCFB

	// Counter
	FlowModeCTR

	// Galois/Counter Mode
	FlowModeGCM

	// Output Feed Back
	FlowModeOFB
)

func (m FlowMode) String() string {
	switch m {
	case FlowModeCBC:
		return "CBC"
	case FlowModeCFB:
		return "CFB"
	case FlowModeCTR:
		return "CTR"
	case FlowModeGCM:
		return "GCM"
	case FlowModeOFB:
		return "OFB"
	}
	return "FlowModeUndefined"
}

////////////////////////////////////////////////////////////////////////////////
