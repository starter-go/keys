package rsad

import (
	"crypto"

	"github.com/starter-go/keys"
)

////////////////////////////////////////////////////////////////////////////////

type publicKeyFacade struct {
	context *publicKeyContext
}

func (inst *publicKeyFacade) _impl() keys.PublicKey {
	return inst
}

func (inst *publicKeyFacade) Class() keys.Class {
	return keys.ClassPublicKey
}

func (inst *publicKeyFacade) Driver() keys.PublicKeyDriver {
	return inst.context.driver
}

func (inst *publicKeyFacade) BaseDriver() keys.Driver {
	return inst.context.driver
}

func (inst *publicKeyFacade) NewEncrypter(opt *keys.Options) (keys.Encrypter, error) {

	if opt == nil {
		opt = new(keys.Options)
	}

	ctx := new(cipherContext)
	ctx.options = *opt
	ctx.private = nil
	ctx.public = inst.context
	ctx.padding = opt.Padding
	ctx.hash = opt.Hash

	ctx.setOptions(opt)
	ctx.encrypter = &encrypter{context: ctx}
	// ctx.decrypter = &decrypter{context: ctx}

	return ctx.encrypter, nil
}

func (inst *publicKeyFacade) NewVerifier(opt *keys.Options) (keys.Verifier, error) {

	if opt == nil {
		opt = new(keys.Options)
	}

	ctx := new(signContext)
	ctx.options = *opt
	ctx.private = nil
	ctx.public = inst.context
	ctx.padding = opt.Padding
	ctx.hash = opt.Hash

	ctx.setOptions(opt)
	ctx.verifier = &verifier{context: ctx}

	return ctx.verifier, nil
}

func (inst *publicKeyFacade) Fingerprint(h crypto.Hash) []byte {
	ls := new(pemLS)
	key := inst.context.raw
	der := ls.getPublicKeyDER(key)
	md := h.New()
	md.Reset()
	md.Write(der)
	return md.Sum(nil)
}

func (inst *publicKeyFacade) Export(want *keys.KeyData) (*keys.KeyData, error) {
	ls := new(pemLS)
	key := inst.context.raw
	return ls.storePublicKey(key)
}

////////////////////////////////////////////////////////////////////////////////

type privateKeyFacade struct {
	context *privateKeyContext
}

func (inst *privateKeyFacade) _impl() keys.PrivateKey {
	return inst
}

func (inst *privateKeyFacade) Class() keys.Class {
	return keys.ClassPrivateKey
}

func (inst *privateKeyFacade) Driver() keys.PrivateKeyDriver {
	return inst.context.driver
}

func (inst *privateKeyFacade) BaseDriver() keys.Driver {
	return inst.context.driver
}

func (inst *privateKeyFacade) Fingerprint(h crypto.Hash) []byte {
	ls := new(pemLS)
	key := &inst.context.raw.PublicKey
	der := ls.getPublicKeyDER(key)
	md := h.New()
	md.Reset()
	md.Write(der)
	return md.Sum(nil)
}

func (inst *privateKeyFacade) Export(want *keys.KeyData) (*keys.KeyData, error) {
	ls := new(pemLS)
	key := inst.context.raw
	return ls.storePrivateKey(key)
}

func (inst *privateKeyFacade) Native() keys.PrivateKeyNative {
	return inst
}

func (inst *privateKeyFacade) Signer() crypto.Signer {
	key := inst.context.raw
	return key
}

func (inst *privateKeyFacade) Decrypter() crypto.Decrypter {
	key := inst.context.raw
	return key
}

func (inst *privateKeyFacade) NewDecrypter(opt *keys.Options) (keys.Decrypter, error) {

	if opt == nil {
		opt = new(keys.Options)
	}

	ctx := new(cipherContext)
	ctx.options = *opt
	ctx.private = inst.context
	ctx.public = inst.context.public
	ctx.padding = opt.Padding
	ctx.hash = opt.Hash

	ctx.setOptions(opt)
	ctx.decrypter = &decrypter{context: ctx}
	// ctx.encrypter = &encrypter{context: ctx}

	return ctx.decrypter, nil
}

func (inst *privateKeyFacade) NewSigner(opt *keys.Options) (keys.Signer, error) {

	if opt == nil {
		opt = new(keys.Options)
	}

	ctx := new(signContext)
	ctx.options = *opt
	ctx.private = inst.context
	ctx.public = inst.context.public
	ctx.padding = opt.Padding
	ctx.hash = opt.Hash

	ctx.setOptions(opt)
	ctx.signer = &signer{context: ctx}

	return ctx.signer, nil
}

func (inst *privateKeyFacade) PublicKey() keys.PublicKey {
	return inst.context.public.facade
}

////////////////////////////////////////////////////////////////////////////////
