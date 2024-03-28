package keys

// PrivateKey ...
type PrivateKey interface {
	Key

	Pair() KeyPair

	NewDecrypter(opt *Options) Decrypter

	NewSigner(opt *Options) Signer
}
