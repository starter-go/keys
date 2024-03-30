package rsad

import (
	"crypto/rsa"
	"fmt"

	"github.com/starter-go/keys"
)

type encrypter struct {
	context *cipherContext
}

func (inst *encrypter) _impl() keys.Encrypter {
	return inst
}

func (inst *encrypter) Key() keys.Key {
	return inst.context.public.facade
}

func (inst *encrypter) Options() keys.Options {
	return inst.context.options
}

func (inst *encrypter) Encrypt(e *keys.Encryption) error {

	hash := inst.context.prepareHash(e)
	key := inst.context.public.raw
	label := e.IV
	mode := inst.context.mode
	plaintext := e.PlainText
	random := inst.context.prepareRandom(e)

	switch mode {
	case CipherModePKCS1v15:
		ciphertext, err := rsa.EncryptPKCS1v15(random, key, plaintext)
		if err != nil {
			return err
		}
		e.CipherText = ciphertext
		break
	case CipherModeSessionKey:
		ciphertext, err := rsa.EncryptPKCS1v15(random, key, plaintext)
		if err != nil {
			return err
		}
		e.CipherText = ciphertext
		break
	case CipherModeOAEP:
		ciphertext, err := rsa.EncryptOAEP(hash, random, key, plaintext, label)
		if err != nil {
			return err
		}
		e.CipherText = ciphertext
		break
	default:
		return fmt.Errorf("bad RSA cipher mode: %d", mode)
	}

	// iv := e.IV
	// cipher.NewCFBEncrypter(key, iv)
	// cipher.NewCFBDecrypter(key, iv)

	return nil
}
