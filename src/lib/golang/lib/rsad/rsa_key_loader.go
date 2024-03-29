package rsad

import "github.com/starter-go/keys"

////////////////////////////////////////////////////////////////////////////////

type privateKeyLoader struct {
	context *rsaContext
}

func (inst *privateKeyLoader) _impl() keys.PrivateKeyLoader {
	return inst
}

func (inst *privateKeyLoader) Load(kd *keys.KeyData) (keys.PrivateKey, error) {

	ls := new(pemLS)
	rawPriKey, err := ls.loadPrivateKey(kd)
	if err != nil {
		return nil, err
	}

	ctx1 := inst.context
	ctx2 := &publicKeyContext{
		parent: ctx1,
		raw:    &rawPriKey.PublicKey,
		driver: ctx1.driverPublic,
	}
	ctx3 := &privateKeyContext{
		parent: ctx1,
		public: ctx2,
		raw:    rawPriKey,
		driver: ctx1.driverPrivate,
	}

	ctx2.facade = &publicKeyFacade{context: ctx2}
	ctx3.facade = &privateKeyFacade{context: ctx3}
	return ctx3.facade, nil
}

////////////////////////////////////////////////////////////////////////////////

type publicKeyLoader struct {
	context *rsaContext
}

func (inst *publicKeyLoader) _impl() keys.PublicKeyLoader {
	return inst
}

func (inst *publicKeyLoader) Load(kd *keys.KeyData) (keys.PublicKey, error) {

	ls := new(pemLS)
	rawPubKey, err := ls.loadPublicKey(kd)
	if err != nil {
		return nil, err
	}

	ctx1 := inst.context
	ctx2 := &publicKeyContext{
		parent: ctx1,
		raw:    rawPubKey,
		driver: ctx1.driverPublic,
	}

	facade := &publicKeyFacade{context: ctx2}
	ctx2.facade = facade
	return facade, nil
}

////////////////////////////////////////////////////////////////////////////////
