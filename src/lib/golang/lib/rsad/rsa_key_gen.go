package rsad

import (
	"crypto/rand"
	"crypto/rsa"

	"github.com/starter-go/keys"
)

type keyGen struct {
	context *rsaContext
}

func (inst *keyGen) _impl() keys.PrivateKeyGenerator {
	return inst
}

func (inst *keyGen) Generate(opt *keys.Options) (keys.PrivateKey, error) {

	// default params

	if opt == nil {
		opt = new(keys.Options)
	}
	bits := opt.Size
	random := opt.Random
	if bits < 1 {
		bits = 1024
	}
	if random == nil {
		random = rand.Reader
	}

	// gen

	key, err := rsa.GenerateKey(random, bits)
	if err != nil {
		return nil, err
	}

	// wrap

	ctx := inst.context
	publicCtx := &publicKeyContext{
		parent: ctx,
		raw:    &key.PublicKey,
		driver: ctx.driverPublic,
		facade: nil,
	}
	privateCtx := &privateKeyContext{
		parent: ctx,
		public: publicCtx,
		raw:    key,
		driver: ctx.driverPrivate,
		facade: nil,
	}

	privateCtx.facade = &privateKeyFacade{context: privateCtx}
	publicCtx.facade = &publicKeyFacade{context: publicCtx}
	return privateCtx.facade, nil
}
