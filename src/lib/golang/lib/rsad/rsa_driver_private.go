package rsad

import "github.com/starter-go/keys"

type rsaPrivateKeyDriver struct {
	context *rsaContext
}

func (inst *rsaPrivateKeyDriver) _impl() keys.PrivateKeyDriver {
	return inst
}

func (inst *rsaPrivateKeyDriver) Algorithm() string {
	return AlgorithmName
}

func (inst *rsaPrivateKeyDriver) Class() keys.Class {
	return keys.ClassPrivateKey
}

func (inst *rsaPrivateKeyDriver) ListRegistrations() []*keys.DriverRegistration {
	return inst.context.driverBase.ListRegistrations()
}

func (inst *rsaPrivateKeyDriver) Loader() keys.PrivateKeyLoader {
	return inst.context.loaderPrivate
}

func (inst *rsaPrivateKeyDriver) Generator() keys.PrivateKeyGenerator {
	return inst.context.keygen
}
