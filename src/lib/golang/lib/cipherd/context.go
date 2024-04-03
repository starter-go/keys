package cipherd

import (
	"crypto/cipher"
	"fmt"

	"github.com/starter-go/keys"
	"github.com/starter-go/keys/paddings"
)

////////////////////////////////////////////////////////////////////////////////

// Context ...
type Context struct {
	// algorithm keys.ComplexAlgorithm

	block   cipher.Block
	padding keys.Padding
	flow    keys.FlowMode
	crypter Crypter
	iv      []byte
}

func (inst *Context) _impl() (keys.Decrypter, keys.Encrypter) {
	return inst, inst
}

func (inst *Context) init(cfg *Config) error {

	// inst.algorithm = cfg.Algorithm
	inst.iv = cfg.IV
	inst.block = cfg.Block
	inst.flow = cfg.Flow
	inst.padding = paddings.NewPadding(cfg.Padding)
	inst.crypter = nil

	return inst.initCrypter()
}

func (inst *Context) initCrypter() error {
	crypter, err := createCrypter(inst)
	if err != nil {
		return err
	}
	inst.crypter = crypter
	return nil
}

// NewBlockMode ...
func (inst *Context) NewBlockMode() (cipher.BlockMode, error) {
	return nil, fmt.Errorf("no impl")
}

// NewAEAD ...
func (inst *Context) NewAEAD() (cipher.AEAD, error) {
	return nil, fmt.Errorf("no impl")
}

// NewStream ...
func (inst *Context) NewStream() (cipher.Stream, error) {
	return nil, fmt.Errorf("no impl")
}

// Encrypt ...
func (inst *Context) Encrypt(e *keys.Crypt) error {
	return inst.crypter.Encrypt(e)
}

// Decrypt ...
func (inst *Context) Decrypt(e *keys.Crypt) error {
	return inst.crypter.Decrypt(e)
}
