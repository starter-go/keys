package cipherd

import (
	"crypto/cipher"
	"io"

	"github.com/starter-go/keys"
)

// Config ...
type Config struct {
	// Algorithm keys.ComplexAlgorithm

	Flow    keys.FlowMode
	Padding keys.PaddingMode
	Block   cipher.Block
	Random  io.Reader
	IV      []byte
}

// NewContext ...
func (inst *Config) NewContext() (*Context, error) {
	ci := new(Context)
	err := ci.init(inst)
	if err != nil {
		return nil, err
	}
	return ci, nil
}
