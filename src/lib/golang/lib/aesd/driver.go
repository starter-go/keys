package aesd

import "github.com/starter-go/keys"

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
	return keys.ClassSecretKey
}

// ListRegistrations ...
func (inst *Driver) ListRegistrations() []*keys.DriverRegistration {

	algorithm := AlgorithmName
	ctx := inst.makeAlgorithmContext()

	// secret-key
	dr2 := &keys.DriverRegistration{
		Algorithm: algorithm,
		Class:     keys.ClassSecretKey,
		Enabled:   true,
		Driver:    ctx.driver,
	}

	// all
	list := make([]*keys.DriverRegistration, 0)
	list = append(list, dr2)
	return list
}

func (inst *Driver) makeAlgorithmContext() *aesContext {

	ctx := &aesContext{}

	ctx.driver = &aesKeyDriver{context: ctx}
	ctx.loader = &aesKeyLoader{context: ctx}
	ctx.keygen = &aesKeyGen{context: ctx}

	return ctx
}

////////////////////////////////////////////////////////////////////////////////

type aesKeyDriver struct {
	context *aesContext
}

func (inst *aesKeyDriver) _impl() keys.SecretKeyDriver {
	return inst
}

// ListRegistrations ...
func (inst *aesKeyDriver) ListRegistrations() []*keys.DriverRegistration {

	// algorithm := AlgorithmName
	// ctx := inst.context

	// secret-key
	// dr2 := &keys.DriverRegistration{
	// 	Algorithm: algorithm,
	// 	Class:     keys.ClassSecretKey,
	// 	Enabled:   true,
	// 	Driver:    ctx.driver,
	// }

	// all
	list := make([]*keys.DriverRegistration, 0)
	// list = append(list, dr2)
	return list
}

// Algorithm ...
func (inst *aesKeyDriver) Algorithm() string {
	return AlgorithmName
}

// Class ...
func (inst *aesKeyDriver) Class() keys.Class {
	return keys.ClassSecretKey
}

func (inst *aesKeyDriver) Loader() keys.SecretKeyLoader {
	return inst.context.loader
}

func (inst *aesKeyDriver) Generator() keys.SecretKeyGenerator {
	return inst.context.keygen
}

////////////////////////////////////////////////////////////////////////////////
