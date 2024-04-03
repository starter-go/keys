package aesd

import (
	"github.com/starter-go/keys"
)

type aesKeyGen struct {
	context *aesContext
}

func (inst *aesKeyGen) _impl() keys.SecretKeyGenerator {
	return inst
}

func (inst *aesKeyGen) Generate(opt *keys.Options) (keys.SecretKey, error) {

	if opt == nil {
		opt = new(keys.Options)
	}

	size := opt.Size

	random := inst.context.getRandom(opt)
	length := size / 8
	parent := inst.context

	buffer := make([]byte, length)
	random.Read(buffer)

	ctx1 := &aesKeyContext{
		parent:     parent,
		rawkey:     buffer,
		sizeInBits: size,
		facade:     nil,
	}

	ctx1.facade = &secretKeyFacade{context: ctx1}

	return ctx1.facade, nil
}
