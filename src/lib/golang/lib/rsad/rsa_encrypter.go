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

	ctx := inst.context
	key := ctx.public.raw
	label := e.IV
	mode := inst.context.padding
	plaintext := e.PlainText
	random := inst.context.prepareRandom()

	switch mode {
	case keys.PaddingPKCS1v15:
		ciphertext, err := rsa.EncryptPKCS1v15(random, key, plaintext)
		if err != nil {
			return err
		}
		e.CipherText = ciphertext
		break
	case keys.PaddingSessionKey:
		ciphertext, err := rsa.EncryptPKCS1v15(random, key, plaintext)
		if err != nil {
			return err
		}
		e.CipherText = ciphertext
		break
	case keys.PaddingOAEP:
		hash := ctx.prepareHashIF()
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
