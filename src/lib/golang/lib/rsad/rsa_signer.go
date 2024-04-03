package rsad

import (
	"crypto/rsa"
	"fmt"

	"github.com/starter-go/keys"
)

type signer struct {
	context *signContext
}

func (inst *signer) _impl() keys.Signer {
	return inst
}

func (inst *signer) Key() keys.Key {
	return inst.context.private.facade
}

func (inst *signer) Options() keys.Options {
	return inst.context.options
}

func (inst *signer) Sign(s *keys.Signature) error {

	if s == nil {
		return fmt.Errorf("param: keys.Signature is nil")
	}

	ctx := inst.context
	key := inst.context.private.raw
	digest := s.Digest
	mode := ctx.padding
	random := ctx.prepareRandom()
	hash := ctx.prepareHashID()

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
