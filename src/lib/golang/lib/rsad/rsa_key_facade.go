package rsad

import (
	"crypto"
	"fmt"

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
	return nil, fmt.Errorf("no impl")
}

func (inst *publicKeyFacade) NewVerifier(opt *keys.Options) (keys.Verifier, error) {
	return nil, fmt.Errorf("no impl")
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
	return nil, fmt.Errorf("no impl")
}

func (inst *privateKeyFacade) NewSigner(opt *keys.Options) (keys.Signer, error) {
	return nil, fmt.Errorf("no impl")
}

func (inst *privateKeyFacade) PublicKey() keys.PublicKey {
	return inst.context.public.facade
}

////////////////////////////////////////////////////////////////////////////////
