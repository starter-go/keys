package rsad

import (
	"crypto/rsa"
	"fmt"

	"github.com/starter-go/keys"
)

type decrypter struct {
	context *cipherContext
}

func (inst *decrypter) _impl() keys.Decrypter {
	return inst
}

func (inst *decrypter) Key() keys.Key {
	return inst.context.private.facade
}

func (inst *decrypter) Options() keys.Options {
	return inst.context.options
}

func (inst *decrypter) Decrypt(e *keys.Encryption) error {

	ciphertext := e.CipherText
	hash := inst.context.prepareHash(e)
	key := inst.context.private.raw
	label := e.IV
	mode := inst.context.mode
	random := inst.context.prepareRandom(e)

	switch mode {
	case CipherModeSessionKey:
		sessionKey := e.PlainText
		err := rsa.DecryptPKCS1v15SessionKey(random, key, ciphertext, sessionKey)
		if err != nil {
			return err
		}
		break
	case CipherModePKCS1v15:
		plaintext, err := rsa.DecryptPKCS1v15(random, key, ciphertext)
		if err != nil {
			return err
		}
		e.PlainText = plaintext
		break
	case CipherModeOAEP:
		plaintext, err := rsa.DecryptOAEP(hash, random, key, ciphertext, label)
		if err != nil {
			return err
		}
		e.PlainText = plaintext
		break
	default:
		return fmt.Errorf("bad RSA decrypt mode: %d", mode)
	}
	return nil
}
