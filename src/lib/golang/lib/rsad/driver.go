package rsad

import "github.com/starter-go/keys"

// AlgorithmName 定义算法名称："RSA"
const AlgorithmName = "RSA"

// Driver ...
type Driver struct {

	//starter:component

	_as func(keys.Driver) //starter:as(".")
}

func (inst *Driver) _impl() keys.Driver {
	return inst
}

// Algorithm ...
func (inst *Driver) Algorithm() string {
	return AlgorithmName
}

// Class ...
func (inst *Driver) Class() keys.Class {
	return keys.ClassPublicKey
}

// ListRegistrations ...
func (inst *Driver) ListRegistrations() []*keys.DriverRegistration {

	algorithm := AlgorithmName
	ctx := inst.makeAlgorithmContext()

	// public-key
	dr2 := &keys.DriverRegistration{
		Algorithm: algorithm,
		Class:     keys.ClassPublicKey,
		Enabled:   true,
		Driver:    ctx.driverPublic,
	}

	// private-key
	dr3 := &keys.DriverRegistration{
		Algorithm: algorithm,
		Class:     keys.ClassPrivateKey,
		Enabled:   true,
		Driver:    ctx.driverPrivate,
	}

	// all
	list := make([]*keys.DriverRegistration, 0)
	list = append(list, dr2)
	list = append(list, dr3)
	return list
}

func (inst *Driver) makeAlgorithmContext() *rsaContext {

	ctx := &rsaContext{}

	ctx.driverBase = inst
	ctx.driverPublic = &rsaPublicKeyDriver{context: ctx}
	ctx.driverPrivate = &rsaPrivateKeyDriver{context: ctx}

	ctx.loaderPrivate = &privateKeyLoader{context: ctx}
	ctx.loaderPublic = &publicKeyLoader{context: ctx}
	ctx.keygen = &keyGen{context: ctx}

	return ctx
}
