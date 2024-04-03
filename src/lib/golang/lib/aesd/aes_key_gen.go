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

	key := make([]byte, length)
	random.Read(key)

	builder := &aesKeyBuilder{
		context: inst.context,
		key:     key,
	}
	return builder.create()
}
