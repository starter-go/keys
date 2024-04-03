package cipherd

import (
	"crypto/cipher"
	"crypto/rand"
	"fmt"

	"github.com/starter-go/keys"
	"github.com/starter-go/keys/paddings"
)

// Crypter ...
type Crypter interface {
	keys.Decrypter
	keys.Encrypter
}

func createCrypter(ctx *Context) (Crypter, error) {

	mode := ctx.flow
	switch mode {

	case keys.FlowModeCBC:
		c := new(crypterCBC)
		c.ch.context = ctx
		return c, nil

	case keys.FlowModeCFB:
		c := new(crypterCFB)
		c.ch.context = ctx
		return c, nil

	case keys.FlowModeCTR:
		c := new(crypterCTR)
		c.ch.context = ctx
		return c, nil

	case keys.FlowModeGCM:
		c := new(crypterGCM)
		c.ch.context = ctx
		return c, nil

	case keys.FlowModeOFB:
		c := new(crypterOFB)
		c.ch.context = ctx
		return c, nil
	}

	return nil, fmt.Errorf("unsupported cipher mode: %s", mode.String())
}

////////////////////////////////////////////////////////////////////////////////

type contextHolder struct {
	context *Context
}

func (inst *contextHolder) prepareBlockIV(e *keys.Encryption) (cipher.Block, []byte, error) {

	block := inst.context.block
	iv := e.IV

	if block == nil {
		return nil, nil, fmt.Errorf("cipher.Block is nil")
	}

	if iv == nil {
		iv = inst.context.iv
	}

	if iv == nil {
		size := block.BlockSize()
		iv = make([]byte, size)
		rand.Reader.Read(iv)
	}

	return block, iv, nil
}

func (inst *contextHolder) getPadding() keys.Padding {
	p := inst.context.padding
	if p == nil {
		p = paddings.NewPKCS7Padding()
		inst.context.padding = p
	}
	return p
}

func (inst *contextHolder) addPadding(blockSize int, data []byte) []byte {
	p := inst.getPadding()
	dst, err := p.AddPadding(data, blockSize)
	if err != nil {
		panic(err)
	}
	return dst
}

func (inst *contextHolder) removePadding(blockSize int, data []byte) []byte {
	p := inst.getPadding()
	dst, err := p.RemovePadding(data, blockSize)
	if err != nil {
		panic(err)
	}
	return dst
}

func (inst *contextHolder) encryptWithBlockMode(bm cipher.BlockMode, e *keys.Encryption) error {
	src := e.PlainText
	src = inst.addPadding(bm.BlockSize(), src)
	dst := make([]byte, len(src))
	bm.CryptBlocks(dst, src)
	e.CipherText = dst
	return nil
}

func (inst *contextHolder) decryptWithBlockMode(bm cipher.BlockMode, e *keys.Encryption) error {
	src := e.CipherText
	dst := make([]byte, len(src))
	bm.CryptBlocks(dst, src)
	dst = inst.removePadding(bm.BlockSize(), dst)
	e.PlainText = dst
	return nil
}

func (inst *contextHolder) encryptWithStream(st cipher.Stream, e *keys.Encryption) error {
	src := e.PlainText
	dst := make([]byte, len(src))
	st.XORKeyStream(dst, src)
	e.CipherText = dst
	return nil
}

func (inst *contextHolder) decryptWithStream(st cipher.Stream, e *keys.Encryption) error {
	src := e.CipherText
	dst := make([]byte, len(src))
	st.XORKeyStream(dst, src)
	e.PlainText = dst
	return nil
}

func (inst *contextHolder) encryptWithAEAD(bm cipher.AEAD, e *keys.Encryption) error {
	return fmt.Errorf("no impl: encryptWithAEAD")
}

func (inst *contextHolder) decryptWithAEAD(bm cipher.AEAD, e *keys.Encryption) error {
	return fmt.Errorf("no impl: decryptWithAEAD")
}

////////////////////////////////////////////////////////////////////////////////

type crypterCBC struct {
	ch contextHolder
}

func (inst *crypterCBC) _impl() Crypter {
	return inst
}

func (inst *crypterCBC) Decrypt(e *keys.Encryption) error {
	b, iv, err := inst.ch.prepareBlockIV(e)
	if err != nil {
		return err
	}
	bm := cipher.NewCBCDecrypter(b, iv)
	return inst.ch.decryptWithBlockMode(bm, e)
}

func (inst *crypterCBC) Encrypt(e *keys.Encryption) error {
	b, iv, err := inst.ch.prepareBlockIV(e)
	if err != nil {
		return err
	}
	bm := cipher.NewCBCEncrypter(b, iv)
	return inst.ch.encryptWithBlockMode(bm, e)
}

////////////////////////////////////////////////////////////////////////////////

type crypterCFB struct {
	ch contextHolder
}

func (inst *crypterCFB) _impl() Crypter {
	return inst
}

func (inst *crypterCFB) Decrypt(e *keys.Encryption) error {
	b, iv, err := inst.ch.prepareBlockIV(e)
	if err != nil {
		return err
	}
	st := cipher.NewCFBDecrypter(b, iv)
	return inst.ch.decryptWithStream(st, e)
}

func (inst *crypterCFB) Encrypt(e *keys.Encryption) error {
	b, iv, err := inst.ch.prepareBlockIV(e)
	if err != nil {
		return err
	}
	st := cipher.NewCFBEncrypter(b, iv)
	return inst.ch.encryptWithStream(st, e)
}

////////////////////////////////////////////////////////////////////////////////

type crypterCTR struct {
	ch contextHolder
}

func (inst *crypterCTR) _impl() Crypter {
	return inst
}

func (inst *crypterCTR) Decrypt(e *keys.Encryption) error {
	b, iv, err := inst.ch.prepareBlockIV(e)
	if err != nil {
		return err
	}
	st := cipher.NewCTR(b, iv)
	return inst.ch.decryptWithStream(st, e)
}

func (inst *crypterCTR) Encrypt(e *keys.Encryption) error {
	b, iv, err := inst.ch.prepareBlockIV(e)
	if err != nil {
		return err
	}
	st := cipher.NewCTR(b, iv)
	return inst.ch.encryptWithStream(st, e)
}

////////////////////////////////////////////////////////////////////////////////

type crypterGCM struct {
	ch contextHolder
}

func (inst *crypterGCM) _impl() Crypter {
	return inst
}

func (inst *crypterGCM) Decrypt(e *keys.Encryption) error {
	b, _, err := inst.ch.prepareBlockIV(e)
	if err != nil {
		return err
	}
	aead, err := cipher.NewGCM(b)
	if err != nil {
		return err
	}
	return inst.ch.decryptWithAEAD(aead, e)
}

func (inst *crypterGCM) Encrypt(e *keys.Encryption) error {
	b, _, err := inst.ch.prepareBlockIV(e)
	if err != nil {
		return err
	}
	aead, err := cipher.NewGCM(b)
	if err != nil {
		return err
	}
	return inst.ch.encryptWithAEAD(aead, e)
}

////////////////////////////////////////////////////////////////////////////////

type crypterOFB struct {
	ch contextHolder
}

func (inst *crypterOFB) _impl() Crypter {
	return inst
}

func (inst *crypterOFB) Decrypt(e *keys.Encryption) error {
	b, iv, err := inst.ch.prepareBlockIV(e)
	if err != nil {
		return err
	}
	st := cipher.NewOFB(b, iv)
	return inst.ch.decryptWithStream(st, e)
}

func (inst *crypterOFB) Encrypt(e *keys.Encryption) error {
	b, iv, err := inst.ch.prepareBlockIV(e)
	if err != nil {
		return err
	}
	st := cipher.NewOFB(b, iv)
	return inst.ch.encryptWithStream(st, e)
}

////////////////////////////////////////////////////////////////////////////////
