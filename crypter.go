package keys

// Encryption 承载加密数据
type Encryption struct {
	CipherText []byte
	PlainText  []byte
	IV         []byte
}

// Decrypter ...
type Decrypter interface {
	Key() Key

	Options() Options

	Decrypt(e *Encryption) error
}

// Encrypter ...
type Encrypter interface {
	Key() Key

	Options() Options

	Encrypt(e *Encryption) error
}
