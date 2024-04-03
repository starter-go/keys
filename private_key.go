package keys

import "crypto"

// PrivateKey ...
type PrivateKey interface {
	Key

	PublicKey() PublicKey

	Decrypter() Decrypter

	Signer() Signer

	Native() PrivateKeyNative

	Driver() PrivateKeyDriver
}

// PrivateKeyNative ...
type PrivateKeyNative interface {
	Signer() crypto.Signer

	Decrypter() crypto.Decrypter
}

// PrivateKeyLoader ... 代表密钥对的 加载器 接口
type PrivateKeyLoader interface {
	Load(kd *KeyData) (PrivateKey, error)
}

// PrivateKeyGenerator ... 代表密钥对的 生成器 接口
type PrivateKeyGenerator interface {
	Generate(opt *Options) (PrivateKey, error)
}

// PrivateKeyDriver 代表密钥对的驱动接口
type PrivateKeyDriver interface {
	Driver

	Loader() PrivateKeyLoader

	Generator() PrivateKeyGenerator
}
