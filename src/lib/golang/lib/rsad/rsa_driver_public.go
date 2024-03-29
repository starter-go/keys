package rsad

import "github.com/starter-go/keys"

type rsaPublicKeyDriver struct {
	context *rsaContext
}

func (inst *rsaPublicKeyDriver) _impl() keys.PublicKeyDriver {
	return inst
}

func (inst *rsaPublicKeyDriver) Algorithm() string {
	return AlgorithmName
}

func (inst *rsaPublicKeyDriver) Class() keys.Class {
	return keys.ClassPublicKey
}

func (inst *rsaPublicKeyDriver) ListRegistrations() []*keys.DriverRegistration {
	return inst.context.driverBase.ListRegistrations()
}

func (inst *rsaPublicKeyDriver) Loader() keys.PublicKeyLoader {
	return inst.context.loaderPublic
}
