package rsad

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"io"

	"github.com/starter-go/keys"
)

type rsaSign struct {
	private *privateKeyContext
	public  *publicKeyContext
}

func (inst *rsaSign) _impl() keys.Signer {
	return inst
}

func (inst *rsaSign) prepareHash(s *keys.Signature) crypto.Hash {
	h := s.Hash
	if h == 0 {
		h = crypto.SHA256
	}
	return h
}

func (inst *rsaSign) preparePadding(s *keys.Signature) keys.PaddingMode {
	mode := s.Padding
	if mode == 0 {
		mode = keys.PaddingPSS
	}
	return mode
}

func (inst *rsaSign) prepareRandom(s *keys.Signature) io.Reader {
	r := s.Random
	if r == nil {
		r = rand.Reader
	}
	return r
}

func (inst *rsaSign) Sign(s *keys.Signature) error {

	if s == nil {
		return fmt.Errorf("param: keys.Signature is nil")
	}

	key := inst.private.raw
	digest := s.Digest
	mode := inst.preparePadding(s)
	random := inst.prepareRandom(s)
	hash := inst.prepareHash(s)

	switch mode {
	case keys.PaddingPKCS1v15:
		sig, err := rsa.SignPKCS1v15(random, key, hash, digest)
		if err != nil {
			return err
		}
		s.Signature = sig
		break
	case keys.PaddingPSS:
		opts := &rsa.PSSOptions{
			Hash:       hash,
			SaltLength: s.SaltLength,
		}
		sig, err := rsa.SignPSS(random, key, hash, digest, opts)
		if err != nil {
			return err
		}
		s.Signature = sig
		break
	default:
		return fmt.Errorf("bad RSA sign mode: %d", mode)
	}
	return nil
}

func (inst *rsaSign) Verify(s *keys.Signature) error {

	if s == nil {
		return fmt.Errorf("param: keys.Signature is nil")
	}

	key := inst.public.raw
	sig := s.Signature
	digest := s.Digest
	mode := inst.preparePadding(s)
	hash := inst.prepareHash(s)

	switch mode {
	case keys.PaddingPKCS1v15:
		return rsa.VerifyPKCS1v15(key, hash, digest, sig)
	case keys.PaddingPSS:
		opts := &rsa.PSSOptions{
			Hash:       hash,
			SaltLength: s.SaltLength,
		}
		return rsa.VerifyPSS(key, hash, digest, sig, opts)
	default:
		break
	}
	return fmt.Errorf("bad RSA sign(verify) mode: %d", mode)
}
