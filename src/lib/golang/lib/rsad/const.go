package rsad

// AlgorithmName 定义算法名称："RSA"
const AlgorithmName = "RSA"

type CipherMode int

const (
	CipherModeOAEP CipherMode = iota + 1
	CipherModePKCS1v15
	CipherModeSessionKey
	CipherModePSS
)
