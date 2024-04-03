package aesd

import "github.com/starter-go/keys"

type aesKeyBuilder struct {
	key     []byte
	context *aesContext
}

func (inst *aesKeyBuilder) create() (keys.SecretKey, error) {

	key := inst.key
	ctx1 := inst.context
	ctx2 := &aesKeyContext{
		parent:     ctx1,
		rawkey:     key,
		sizeInBits: len(key) * 8,
		facade:     nil,
		enc:        nil,
		dec:        nil,
	}

	ci := &aesCipherImpl{context: ctx2}
	facade := &secretKeyFacade{context: ctx2}

	ctx2.facade = facade
	ctx2.dec = ci
	ctx2.enc = ci

	return ctx2.facade, nil
}
