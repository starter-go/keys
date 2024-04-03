package keys

import (
	"crypto"
	"io"
)

// Signature ...
type Signature struct {

	// options

	Random     io.Reader
	SaltLength int
	Hash       crypto.Hash
	Padding    PaddingMode

	// input/output

	Digest    []byte
	Signature []byte
}

// Signer ...
type Signer interface {
	Sign(s *Signature) error
}

// Verifier 代表签名验证器
type Verifier interface {
	Verify(s *Signature) error
}
