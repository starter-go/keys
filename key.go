package keys

import (
	"crypto"
	"io"
)

// Class 表示密钥类型
type Class string

// 定义各种密钥类型
const (
	ClassPrivateKey Class = "private" // 私钥
	ClassPublicKey  Class = "public"  // 公钥
	ClassSecretKey  Class = "secret"  // 对称密钥
)

// Key ...
type Key interface {
	BaseDriver() Driver

	Class() Class

	Export(want *KeyData) (*KeyData, error)

	Fingerprint(h crypto.Hash) []byte
}

// KeyData 包含密钥的序列化数据
type KeyData struct {
	Algorithm   string // like 'rsa'
	Encoding    string // like 'x509'
	Format      string // like 'pem'
	ContentType string // like 'application/x-pem'
	Content     []byte
}

// Options 是通用的选项
type Options struct {
	Algorithm ComplexAlgorithm
	Size      int // in bits
	Random    io.Reader
}
