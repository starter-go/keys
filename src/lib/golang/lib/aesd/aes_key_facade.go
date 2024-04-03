package aesd

import (
	"crypto"
	"crypto/aes"
	"crypto/cipher"

	"github.com/starter-go/base/lang"
	"github.com/starter-go/keys"
)

type secretKeyFacade struct {
	context *aesKeyContext
}

func (inst *secretKeyFacade) _impl() keys.SecretKey {
	return inst
}

func (inst *secretKeyFacade) Algorithm() string {
	return AlgorithmName
}

func (inst *secretKeyFacade) Class() keys.Class {
	return keys.ClassSecretKey
}

func (inst *secretKeyFacade) Driver() keys.SecretKeyDriver {
	return inst.context.parent.driver
}

func (inst *secretKeyFacade) BaseDriver() keys.Driver {
	return inst.context.parent.driver
}

func (inst *secretKeyFacade) createCipher(opt *keys.Options) (aesCipher, error) {
	cipher1 := &aesCipherImpl{
		context: inst.context,
	}
	// cipher1.init(opt)
	return cipher1, nil
}

func (inst *secretKeyFacade) Decrypter() keys.Decrypter {
	return inst.context.dec
}

func (inst *secretKeyFacade) Encrypter() keys.Encrypter {
	return inst.context.enc
}

func (inst *secretKeyFacade) Native() keys.SecretKeyNative {
	return inst
}

func (inst *secretKeyFacade) Key() keys.SecretKey {
	return inst
}

func (inst *secretKeyFacade) NewCipher() (cipher.Block, error) {
	key := inst.context.rawkey
	return aes.NewCipher(key)
}

func (inst *secretKeyFacade) Export(want *keys.KeyData) (*keys.KeyData, error) {

	key := inst.context.rawkey
	b64 := lang.Base64FromBytes(key)

	result := new(keys.KeyData)
	result.Algorithm = AlgorithmName
	result.Content = []byte(b64)
	result.ContentType = "application/x-aes-key-base64"
	result.Encoding = "base64"
	result.Format = "base64"

	return result, nil
}

func (inst *secretKeyFacade) Fingerprint(h crypto.Hash) []byte {
	key := inst.context.rawkey
	md := h.New()
	md.Write(key)
	sum := md.Sum(nil)
	return sum
}
