package aesd

import (
	"github.com/starter-go/base/lang"
	"github.com/starter-go/keys"
)

type aesKeyLoader struct {
	context *aesContext
}

func (inst *aesKeyLoader) _impl() keys.SecretKeyLoader {
	return inst
}

func (inst *aesKeyLoader) Load(kd *keys.KeyData) (keys.SecretKey, error) {

	key, err := inst.loadBytes(kd)
	if err != nil {
		return nil, err
	}

	ctx1 := inst.context
	ctx2 := &aesKeyContext{
		parent:     ctx1,
		rawkey:     key,
		sizeInBits: len(key) * 8,
		facade:     nil,
		block0:     nil,
	}

	ctx2.facade = &secretKeyFacade{context: ctx2}

	return ctx2.facade, nil
}

func (inst *aesKeyLoader) loadBytes(kd *keys.KeyData) ([]byte, error) {
	b64 := lang.Base64(kd.Content)
	return b64.Bytes(), nil
}
