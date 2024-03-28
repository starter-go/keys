package keys

import "crypto"

// KeyPair ... (aka. private_key)
type KeyPair interface {
	Key

	KeyPairDriver() KeyPairDriver

	PublicKey() PublicKey

	PrivateKey() PrivateKey

	Native() KeyPairNative
}

// KeyPairNative ...
type KeyPairNative interface {
	Signer() crypto.Signer

	Decrypter() crypto.Decrypter
}

// KeyPairLoader ... 代表密钥对的 加载器 接口
type KeyPairLoader interface {
	Load(kd *KeyData) (KeyPair, error)
}

// KeyPairGenerator ... 代表密钥对的 生成器 接口
type KeyPairGenerator interface {
	Generate(opt *Options) (KeyPair, error)
}

// KeyPairDriver 代表密钥对的驱动接口
type KeyPairDriver interface {
	Driver

	Loader() KeyPairLoader

	Generator() KeyPairGenerator
}
