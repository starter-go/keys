package keys

// PublicKey ...
type PublicKey interface {
	Key

	Driver() PublicKeyDriver

	NewEncrypter(opt *Options) (Encrypter, error)

	NewVerifier(opt *Options) (Verifier, error)
}

// PublicKeyLoader ... 代表公钥的 加载器 接口
type PublicKeyLoader interface {
	Load(kd *KeyData) (PublicKey, error)
}

// PublicKeyDriver 代表公钥的驱动接口
type PublicKeyDriver interface {
	Driver

	Loader() PublicKeyLoader
}
