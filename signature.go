package keys

import "io"

// Signature ...
type Signature struct {
	Rand      io.Reader
	Digest    []byte
	Signature []byte
}

// Signer ...
type Signer interface {
	Key() Key

	Options() Options

	Sign(s *Signature) error
}

// Verifier 代表签名验证器
type Verifier interface {
	Key() Key

	Options() Options

	Verify(s *Signature) error
}
