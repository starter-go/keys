package aesd

import (
	"crypto/cipher"
	"crypto/rand"
	"io"

	"github.com/starter-go/keys"
)

type aesContext struct {
	driver keys.SecretKeyDriver
	loader keys.SecretKeyLoader
	keygen keys.SecretKeyGenerator
}

type aesKeyContext struct {
	parent     *aesContext
	rawkey     []byte
	sizeInBits int
	facade     keys.SecretKey
	block0     cipher.Block // the block as prototype
}

// type aesOptionContext struct {
// 	parent      *aesKeyContext
// 	iv          []byte
// 	mode        keys.CipherMode
// 	paddingMode keys.PaddingMode
// 	method      keys.CipherMethod
// 	options     keys.Options

// 	enc keys.Encrypter
// 	dec keys.Decrypter
// }

////////////////////////////////////////////////////////////////////////////////

func (inst *aesContext) getRandom(opt *keys.Options) io.Reader {
	if opt != nil {
		if opt.Random != nil {
			return opt.Random
		}
	}
	return rand.Reader
}

////////////////////////////////////////////////////////////////////////////////

// func (inst *aesOptionContext) clone() *aesOptionContext {
// 	src := inst
// 	dst := new(aesOptionContext)
// 	*dst = *src
// 	return dst
// }

// func (inst *aesOptionContext) loadFromOption(opt *keys.Options) {
// }

// func (inst *aesOptionContext) prepareBlock() (cipher.Block, error) {

// }

// func (inst *aesOptionContext) preparePadding() (keys.Padding, error) {

// }

// func (inst *aesOptionContext) prepareIV(e *keys.Encryption) ([]byte, error) {

// }

// func (inst *aesOptionContext) prepareBlockPaddingIV(e *keys.Encryption) (block cipher.Block, padding keys.Padding, iv []byte, err error) {

// 	block, err = inst.prepareBlock()
// 	if err != nil {
// 		return nil, nil, nil, err
// 	}

// 	padding, err = inst.preparePadding()
// 	if err != nil {
// 		return nil, nil, nil, err
// 	}

// 	iv, err = inst.prepareIV(e)
// 	if err != nil {
// 		return nil, nil, nil, err
// 	}

// 	return
// }

////////////////////////////////////////////////////////////////////////////////
