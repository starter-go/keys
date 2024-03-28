package keys

import "io"

// Encryption 承载加密数据
type Encryption struct {
	Rand       io.Reader
	CipherText []byte
	PlainText  []byte
	IV         []byte
}

// Decrypter ...
type Decrypter interface {
	Key() Key

	Options() Options

	Encrypt(e *Encryption) error
}

// Encrypter ...
type Encrypter interface {
	Key() Key

	Options() Options

	Encrypt(e *Encryption) error
}
