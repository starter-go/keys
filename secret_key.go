package keys

import (
	"crypto/cipher"
)

// SecretKey ...
type SecretKey interface {
	Key

	Driver() SecretKeyDriver

	Decrypter() Decrypter

	Encrypter() Encrypter

	Native() SecretKeyNative
}

// SecretKeyNative ...
type SecretKeyNative interface {
	Key() SecretKey

	NewCipher() (cipher.Block, error)
}

// SecretKeyLoader ... 代表密钥的 加载器 接口
type SecretKeyLoader interface {
	Load(kd *KeyData) (SecretKey, error)
}

// SecretKeyGenerator ... 代表密钥的 生成器 接口
type SecretKeyGenerator interface {
	Generate(opt *Options) (SecretKey, error)
}

// SecretKeyDriver 代表对称密钥的驱动接口
type SecretKeyDriver interface {
	Driver

	Loader() SecretKeyLoader

	Generator() SecretKeyGenerator
}
