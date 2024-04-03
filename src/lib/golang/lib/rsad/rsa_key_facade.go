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

func (inst *publicKeyFacade) Encrypter() keys.Encrypter {
	return inst
}

func (inst *publicKeyFacade) Verifier() keys.Verifier {
	return inst
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

func (inst *publicKeyFacade) Verify(s *keys.Signature) error {
	si := &rsaSign{
		public: inst.context,
	}
	return si.Verify(s)
}

func (inst *publicKeyFacade) Encrypt(e *keys.Crypt) error {
	ci := &rsaCipher{
		public: inst.context,
	}
	return ci.Encrypt(e)
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
	n := &privateKeyNative{
		context: inst.context,
	}
	return n
}

func (inst *privateKeyFacade) Decrypter() keys.Decrypter {
	return inst
}

func (inst *privateKeyFacade) Signer() keys.Signer {
	return inst
}

func (inst *privateKeyFacade) PublicKey() keys.PublicKey {
	return inst.context.public.facade
}

func (inst *privateKeyFacade) Sign(s *keys.Signature) error {
	si := &rsaSign{
		public:  inst.context.public,
		private: inst.context,
	}
	return si.Sign(s)
}

func (inst *privateKeyFacade) Decrypt(e *keys.Crypt) error {
	ci := &rsaCipher{
		public:  inst.context.public,
		private: inst.context,
	}
	return ci.Decrypt(e)
}

////////////////////////////////////////////////////////////////////////////////

type privateKeyNative struct {
	context *privateKeyContext
}

func (inst *privateKeyNative) _impl() keys.PrivateKeyNative {
	return inst
}

func (inst *privateKeyNative) Decrypter() crypto.Decrypter {
	key := inst.context.raw
	return key
}

func (inst *privateKeyNative) Signer() crypto.Signer {
	key := inst.context.raw
	return key
}
