package rsad

import (
	"crypto/rsa"
	"fmt"

	"github.com/starter-go/keys"
)

type verifier struct {
	context *signContext
}

func (inst *verifier) _impl() keys.Verifier {
	return inst
}

func (inst *verifier) Key() keys.Key {
	return inst.context.public.facade
}

func (inst *verifier) Options() keys.Options {
	return inst.context.options
}

func (inst *verifier) Verify(s *keys.Signature) error {

	if s == nil {
		return fmt.Errorf("param: keys.Signature is nil")
	}

	ctx := inst.context
	key := inst.context.public.raw
	digest := s.Digest
	mode := ctx.padding
	sig := s.Signature
	hash := ctx.prepareHashID()

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
