package rsad

import (
	"crypto"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"hash"
	"io"

	"github.com/starter-go/keys"
)

type rsaCipher struct {
	private *privateKeyContext
	public  *publicKeyContext
}

func (inst *rsaCipher) _impl() (keys.Encrypter, keys.Decrypter) {
	return inst, inst
}

func (inst *rsaCipher) prepareHash(e *keys.Crypt) hash.Hash {
	h := e.Hash
	if h == 0 {
		h = crypto.SHA256
	}
	return h.New()
}

func (inst *rsaCipher) preparePadding(e *keys.Crypt) keys.PaddingMode {
	mode := e.Padding
	if mode == 0 {
		mode = keys.PaddingOAEP
	}
	return mode
}

func (inst *rsaCipher) prepareRandom(e *keys.Crypt) io.Reader {
	r := e.Random
	if r == nil {
		r = rand.Reader
	}
	return r
}

func (inst *rsaCipher) Encrypt(e *keys.Crypt) error {

	ctx := inst.public
	key := ctx.raw
	label := e.IV
	plaintext := e.PlainText
	mode := inst.preparePadding(e)
	random := inst.prepareRandom(e)

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
		hash := inst.prepareHash(e)
		ciphertext, err := rsa.EncryptOAEP(hash, random, key, plaintext, label)
		if err != nil {
			return err
		}
		e.CipherText = ciphertext
		break
	default:
		return fmt.Errorf("bad RSA cipher mode: %d", mode)
	}

	return nil
}

func (inst *rsaCipher) Decrypt(e *keys.Crypt) error {

	ctx := inst.private
	key := ctx.raw
	ciphertext := e.CipherText
	label := e.IV
	mode := inst.preparePadding(e)
	random := inst.prepareRandom(e)

	switch mode {
	case keys.PaddingSessionKey:
		sessionKey := e.PlainText
		err := rsa.DecryptPKCS1v15SessionKey(random, key, ciphertext, sessionKey)
		if err != nil {
			return err
		}
		break
	case keys.PaddingPKCS1v15:
		plaintext, err := rsa.DecryptPKCS1v15(random, key, ciphertext)
		if err != nil {
			return err
		}
		e.PlainText = plaintext
		break
	case keys.PaddingOAEP:
		hash := inst.prepareHash(e)
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

func (inst *rsaCipher) Block() cipher.Block {
	panic("unsupported method:Block()")
}
