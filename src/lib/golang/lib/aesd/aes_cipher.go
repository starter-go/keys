package aesd

import (
	"crypto/aes"
	"crypto/cipher"

	"github.com/starter-go/keys"
	"github.com/starter-go/keys/src/lib/golang/lib/cipherd"
)

type aesCipher interface {
	keys.Encrypter
	keys.Decrypter
}

////////////////////////////////////////////////////////////////////////////////

type aesCipherImpl struct {
	context *aesKeyContext

	// algorithm keys.ComplexAlgorithm
	// random  io.Reader
	// flow    keys.FlowMode
	// padding keys.PaddingMode
	// iv      []byte

}

func (inst *aesCipherImpl) _impl() aesCipher {
	return inst
}

func (inst *aesCipherImpl) Block() cipher.Block {
	key := inst.context.rawkey
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	return block
}

// func (inst *aesCipherImpl) init(opt *keys.Options) error {

// 	inst.flow = opt.Flow
// 	inst.padding = opt.Padding
// 	inst.iv = opt.IV
// 	inst.random = opt.Random

// 	return nil
// }

func (inst *aesCipherImpl) config(e *keys.Crypt) (*cipherd.Config, error) {

	ctx := inst.context
	key := ctx.rawkey
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	cfg := &cipherd.Config{
		// Algorithm: inst.algorithm,
		Flow:    e.Flow,
		Padding: e.Padding,
		IV:      e.IV,
		Random:  e.Random,
		Block:   block,
	}

	return cfg, nil
}

func (inst *aesCipherImpl) Decrypt(e *keys.Crypt) error {
	cfg, err := inst.config(e)
	if err != nil {
		return err
	}
	ctx, err := cfg.NewContext()
	if err != nil {
		return err
	}
	return ctx.Decrypt(e)
}

func (inst *aesCipherImpl) Encrypt(e *keys.Crypt) error {
	cfg, err := inst.config(e)
	if err != nil {
		return err
	}
	ctx, err := cfg.NewContext()
	if err != nil {
		return err
	}
	return ctx.Encrypt(e)
}
