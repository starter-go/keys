package keys

import (
	"crypto"
	"crypto/cipher"
	"io"
)

// Crypt 承载加密数据
type Crypt struct {
	Flow    FlowMode    // 数据块串流模式
	Padding PaddingMode // 填充模式
	Random  io.Reader   // 随机数生成器
	IV      []byte      // 初始化向量
	Hash    crypto.Hash // 摘要算法

	CipherText []byte // 密文
	PlainText  []byte // 明文
}

// Decrypter ...
type Decrypter interface {
	Decrypt(e *Crypt) error
	Block() cipher.Block
}

// Encrypter ...
type Encrypter interface {
	Block() cipher.Block
	Encrypt(e *Crypt) error
}
